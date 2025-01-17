package Redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"time"
)

var ErrInvalidURL = errors.New("invalid URL")

const (
	scope  = "internal.cache.Redis."
	docTTL = 15 * time.Minute
)

type Redis struct {
	cl *redis.Client
}

func New(addr, pass string, db int) (*Redis, error) {
	const op = scope + "New"
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	if err := rdb.Ping(context.TODO()).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Redis{cl: rdb}, nil
}

func (r *Redis) Create(ctx context.Context, doc *docsv1.UserTDocument) error {
	const op = scope + "Create"

	if err := r.cl.Exists(ctx, doc.Url).Err(); errors.Is(err, redis.Nil) {
		return fmt.Errorf("%s: %w", op, err)
	}

	now := uint64(time.Now().UnixNano())
	if err := r.cl.Set(ctx, doc.Url, &docsv1.TDocument{
		Url:            doc.Url,
		PubDate:        now,
		FetchTime:      now,
		Text:           doc.Text,
		FirstFetchTime: now,
	}, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *Redis) Get(ctx context.Context, url string) (*docsv1.TDocument, error) {
	const op = scope + "Get"

	var tdoc docsv1.TDocument
	if err := r.cl.Get(ctx, url).Scan(&tdoc); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &tdoc, nil
}

func fixDoc(doc *docsv1.TDocument) {
	now := uint64(time.Now().UnixNano())
	if doc.FirstFetchTime == 0 {
		doc.FirstFetchTime = now
	}
	if doc.FetchTime == 0 {
		doc.FetchTime = now
	}
	if doc.PubDate == 0 {
		doc.PubDate = now
	}

	if doc.FetchTime < doc.FirstFetchTime {
		doc.FirstFetchTime = doc.FetchTime
	}
}

func (r *Redis) Process(ctx context.Context, doc *docsv1.TDocument) (*docsv1.TDocument, error) {
	const op = scope + "Save"

	var savedDoc docsv1.TDocument
	if err := r.cl.Get(ctx, doc.Url).Scan(&savedDoc); errors.Is(err, redis.Nil) {
		fixDoc(doc)
		if err = r.saveTDoc(ctx, doc); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		return doc, nil
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	res := compareAndEdit(doc, &savedDoc)

	if err := r.cl.Set(ctx, doc.Url, res, docTTL).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}

func (r *Redis) Update(ctx context.Context, doc *docsv1.UserTDocument) error {
	const op = scope + "Update"

	if len(doc.Url) == 0 {
		return ErrInvalidURL
	}

	var tdoc docsv1.TDocument

	if err := r.cl.Get(ctx, doc.Url).Scan(&tdoc); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	tdoc.Text = doc.Text
	tdoc.FetchTime = uint64(time.Now().UnixNano())

	if err := r.cl.Set(ctx, doc.Url, &tdoc, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Redis) saveTDoc(ctx context.Context, doc *docsv1.TDocument) error {
	const op = scope + "saveTDoc"

	if err := r.cl.Set(ctx, doc.Url, doc, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func compareAndEdit(doc1 *docsv1.TDocument, doc2 *docsv1.TDocument) *docsv1.TDocument {
	var res = docsv1.TDocument{
		Url:            doc1.Url,
		PubDate:        doc1.PubDate,
		FetchTime:      doc1.FetchTime,
		Text:           doc1.Text,
		FirstFetchTime: doc1.FirstFetchTime,
	}

	if doc1.FetchTime == 0 && doc2.FetchTime == 0 {
		res.FetchTime = uint64(time.Now().UnixNano())
	} else if doc1.FetchTime == 0 {
		res.FetchTime = doc2.FetchTime
	} else if doc2.FetchTime == 0 {
		res.FetchTime = doc1.FetchTime
	} else if doc1.FetchTime < doc2.FetchTime {
		res.FetchTime = doc2.FetchTime
		res.Text = doc2.Text
	} else {
		res.PubDate = doc2.PubDate
	}

	if doc1.FirstFetchTime == 0 && doc2.FirstFetchTime == 0 {
		res.FirstFetchTime = uint64(time.Now().UnixNano())
	} else if doc1.FirstFetchTime == 0 {
		res.FirstFetchTime = doc2.FirstFetchTime
	} else if doc2.FirstFetchTime == 0 {
		res.FirstFetchTime = doc1.FirstFetchTime
	} else if doc1.FirstFetchTime > doc2.FirstFetchTime {
		res.FirstFetchTime = doc2.FirstFetchTime
	} else {
		res.FirstFetchTime = doc1.FirstFetchTime
	}

	if res.FetchTime < res.FirstFetchTime {
		res.FirstFetchTime = res.FetchTime
	} else if doc2.FetchTime < res.FirstFetchTime {
		res.FirstFetchTime = doc2.FetchTime
	} else if doc1.FetchTime < res.FirstFetchTime {
		res.FirstFetchTime = doc1.FetchTime
	}

	return &res
}
