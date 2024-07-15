package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/vk_test/internal/lib/sl"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"google.golang.org/protobuf/proto"
	"log/slog"
	"time"
)

const (
	scope   = "internal.broker.nats."
	receive = "docs"
	send    = "tdocs"
	qrcv    = "rcv"
	qsnd    = "snd"

	reqTTL = 10 * time.Second
)

type Storage interface {
	Process(ctx context.Context, tdoc *docsv1.TDocument) (*docsv1.TDocument, error)
}

type Nats struct {
	ns  *nats.Conn
	s   Storage
	log *slog.Logger
}

func New(log *slog.Logger, s Storage, address string, retry bool, maxReconn int, reconnWait time.Duration) (*Nats, error) {
	const op = scope + "New"

	nc, err := nats.Connect(address,
		nats.RetryOnFailedConnect(retry),
		nats.MaxReconnects(maxReconn),
		nats.ReconnectWait(reconnWait),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Nats{ns: nc, log: log, s: s}, nil
}

func (n *Nats) Process(ctx context.Context, tdoc *docsv1.TDocument) (*docsv1.TDocument, error) {
	const op = scope + "Process"

	data, err := proto.Marshal(tdoc)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	msg, err := n.ns.Request(receive, data, reqTTL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var resp docsv1.TDocument
	if err = proto.Unmarshal(msg.Data, &resp); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &resp, nil
}

func (n *Nats) Start(ctx context.Context) error {
	const op = scope + "Start"

	sub, err := n.ns.QueueSubscribe(receive, qrcv, func(msg *nats.Msg) {
		// uncomment for real case. Publish to reworked data to queue
		var tdoc docsv1.TDocument
		if err := proto.Unmarshal(msg.Data, &tdoc); err != nil {
			n.log.Error("couldn't unmarshal TDocument", sl.Op(op), sl.Err(err))
			return
		}

		tdocNew, err := n.s.Process(ctx, &tdoc)
		if err != nil {
			slog.Error("couldn't process TDocument", sl.Op(op), sl.Err(err))
			return
		}

		data, err := proto.Marshal(tdocNew)
		if err != nil {
			n.log.Error("couldn't get TDocument", sl.Op(op), sl.Err(err))
			return
		}

		if err = n.ns.Publish(send, data); err != nil {
			n.log.Error("couldn't publish message to <send> queue", sl.Op(op), sl.Err(err))
			return
		}
		// needed for test
		if err = msg.Respond(data); err != nil {
			n.log.Error("couldn't publish response", sl.Op(op), sl.Err(err))
			return
		}
	})
	if err != nil {
		return fmt.Errorf("%s: %s: %w", op, receive, err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("%s: %w", op, err)
		}
	}
}
