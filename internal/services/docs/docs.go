package docs

import (
	"context"
	"fmt"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
)

type Broker interface {
	Process(ctx context.Context, document *docsv1.TDocument) (*docsv1.TDocument, error)
}

type Storage interface {
	Create(ctx context.Context, doc *docsv1.UserTDocument) error
	Get(ctx context.Context, url string) (*docsv1.TDocument, error)
	Update(ctx context.Context, doc *docsv1.UserTDocument) error
}

const scope = "internal.services.docsGRPC."

type Handler struct {
	broker  Broker
	storage Storage
}

func New(b Broker, s Storage) *Handler {
	return &Handler{
		broker:  b,
		storage: s,
	}
}

// Create -- func for endpoint case
func (h *Handler) Create(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	const op = scope + "Create"

	if err := h.storage.Create(ctx, document); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tdoc, err := h.storage.Get(ctx, document.Url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tdoc, nil
}

// Update -- func for endpoint case
func (h *Handler) Update(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	const op = scope + "Update"

	if err := h.storage.Update(ctx, document); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tdoc, err := h.storage.Get(ctx, document.Url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tdoc, nil
}

// Process -- func for endpoint
func (h *Handler) Process(ctx context.Context, document *docsv1.TDocument) (*docsv1.TDocument, error) {
	const op = scope + "Process"

	tdoc, err := h.broker.Process(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tdoc, nil
}

// Get -- func for endpoint case
func (h *Handler) Get(ctx context.Context, url string) (*docsv1.TDocument, error) {
	const op = scope + "Get"

	tdoc, err := h.storage.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tdoc, nil
}
