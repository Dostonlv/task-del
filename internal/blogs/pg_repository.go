//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package blogs

import (
	"context"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/pkg/utils"

	"github.com/google/uuid"
)

// Repository Blogs repository interface
type Repository interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Update(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, blogID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error)
}
