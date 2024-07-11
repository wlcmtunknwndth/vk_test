package docs

import (
	"context"
	"github.com/wlcmtunknwndth/vk_test/internal/domain/models"
)

type Broker interface {
	Create(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error)
	Update(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error)
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
	broker Broker
	cache  Storage
}

func (h *Handler) Create(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error) {
	const op = scope + "Create"

}

func (h *Handler) Update(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error) {
	const op = scope + "Update"
}

func (h *Handler) Process(ctx context.Context, document *models.TDocument) (*models.TDocument, error) {
	const op = scope + "Process"

}

func (h *Handler) Get(ctx context.Context, url string) (*models.TDocument, error) {

}
