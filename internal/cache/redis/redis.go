package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/wlcmtunknwndth/vk_test/internal/domain/models"
	"time"
)

//Create(ctx context.Context, doc *models.UserTDocument) error
//Get(ctx context.Context, url string) (*models.TDocument, error)
//Save(ctx context.Context, doc *models.UserTDocument) error

var ErrInvalidURL = errors.New("invalid URL")

const (
	scope  = "internal.cache.redis."
	docTTL = 15 * time.Minute
)

type Redis struct {
	cl *redis.Client
}

func New(addr, pass string, db int) *Redis {
	const op = scope + "New"

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return &Redis{cl: rdb}
}

func (r *Redis) Create(ctx context.Context, doc *models.UserTDocument) error {
	const op = scope + "Create"

	if err := r.cl.Exists(ctx, doc.Url).Err(); errors.Is(err, redis.Nil) {
		return fmt.Errorf("%s: %w", op, err)
	}

	now := time.Now().UnixNano().Uint64()
	if err := r.cl.Set(ctx, doc.Url, &models.TDocument{
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

func (r *Redis) Get(ctx context.Context, url string) (*models.TDocument, error) {
	const op = scope + "Get"

	var tdoc models.TDocument
	if err := r.cl.Get(ctx, url).Scan(&tdoc); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &tdoc, nil
}

func (r *Redis) Save(ctx context.Context, doc *models.TDocument) error {
	const op = scope + "Save"

	var savedDoc models.TDocument
	if err := r.cl.Get(ctx, doc.Url).Scan(&savedDoc); errors.Is(err, redis.Nil) {
		if err = r.saveTDoc(ctx, doc); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	} else if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	//if savedDoc.FetchTime < doc.FetchTime {
	//	savedDoc.FetchTime = doc.FetchTime
	//	savedDoc.Text = doc.Text
	//} else {
	//	savedDoc.PubDate = doc.PubDate
	//}
	//
	//if savedDoc.FirstFetchTime > doc.FirstFetchTime {
	//	savedDoc.FirstFetchTime = doc.FirstFetchTime
	//}
	res := compareAndEdit(&savedDoc, doc)

	if err := r.cl.Set(ctx, doc.Url, res, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Redis) Update(ctx context.Context, doc *models.UserTDocument) error {
	const op = scope + "Update"

	if len(doc.Url) == 0 {
		return ErrInvalidURL
	}

	var tdoc models.TDocument

	if err := r.cl.Get(ctx, doc.Url).Scan(&tdoc); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	tdoc.Text = doc.Text
	tdoc.FetchTime = time.Now().UnixNano().Uint64()

	if err := r.cl.Set(ctx, doc.Url, tdoc, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Redis) saveTDoc(ctx context.Context, doc *models.TDocument) error {
	const op = scope + "saveTDoc"

	if err := r.cl.Set(ctx, doc.Url, doc, docTTL).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func compareAndEdit(doc1 *models.TDocument, doc2 *models.TDocument) *models.TDocument {
	var res = *doc1

	if doc1.FetchTime < doc2.FetchTime {
		res.FetchTime = doc2.FetchTime
		res.Text = doc2.Text
	} else {
		res.PubDate = doc2.PubDate
	}

	if doc1.FirstFetchTime > doc2.FirstFetchTime {
		res.FirstFetchTime = doc2.FirstFetchTime
	}

	return &res
}
