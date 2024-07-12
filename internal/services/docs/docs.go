package docs

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/vk_test/internal/domain/models"
)

type Broker interface {
	Process(ctx context.Context, document *models.TDocument) (*models.TDocument, error)
}

type Storage interface {
	Create(ctx context.Context, doc *models.UserTDocument) error
	Get(ctx context.Context, url string) (*models.TDocument, error)
	Save(ctx context.Context, doc *models.TDocument) error
	Update(ctx context.Context, doc *models.UserTDocument) error
}

const scope = "internal.services.docs."

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
func (h *Handler) Create(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error) {
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
func (h *Handler) Update(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error) {
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

// Process -- func for use as middleware
func (h *Handler) Process(ctx context.Context, document *models.TDocument) (*models.TDocument, error) {
	const op = scope + "Process"

	panic(op + ": implement me!!!")
}

// Get -- func for endpoint case
func (h *Handler) Get(ctx context.Context, url string) (*models.TDocument, error) {
	const op = scope + "Get"

	tdoc, err := h.storage.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return tdoc, nil
}
