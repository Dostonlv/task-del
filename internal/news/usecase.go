package news

import (
	"context"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/pkg/utils"
	"github.com/google/uuid"
)

// news use case interface
type UseCase interface {
	Create(ctx context.Context, news *models.New) (*models.New, error)
	Update(ctx context.Context, news *models.New) (*models.New, error)
	Delete(ctx context.Context, newsID uuid.UUID) error
	GetByID(ctx context.Context, newsID uuid.UUID) (*models.New, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
