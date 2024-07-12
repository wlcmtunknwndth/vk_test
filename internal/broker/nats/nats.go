package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/vk_test/internal/domain/models"
)

const scope = "internal.broker.nats."

type Nats struct {
	ns *nats.Conn
}

func New(urlSender, urlReceiver string, opts ...nats.Option) (*Nats, error) {
	const op = scope + "New"

	nc, err := nats.Connect(urlSender, opts...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Nats{ns: nc}, nil
}

func (n *Nats) Process(ctx context.Context, tdoc *models.TDocument) (*models.TDocument, error) {
	const op = scope + "Process"

	panic(op + ": implement me!!!")
}
