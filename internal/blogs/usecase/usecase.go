package usecase

import (
	"context"
	"github.com/Dostonlv/task-del/config"
	"github.com/Dostonlv/task-del/internal/blogs"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/Dostonlv/task-del/pkg/utils"

	"github.com/google/uuid"
)

// blogs UseCase
type blogsUC struct {
	cfg       *config.Config
	blogsRepo blogs.Repository
	logger    logger.Logger
}

// NewBlogsUseCase Blogs UseCase constructor
func NewBlogsUseCase(cfg *config.Config, blogsRepo blogs.Repository, logger logger.Logger) blogs.UseCase {
	return &blogsUC{cfg: cfg, blogsRepo: blogsRepo, logger: logger}
}

// Create blog
func (u *blogsUC) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	return u.blogsRepo.Create(ctx, blog)
}

// Update blog
func (u *blogsUC) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	updatedBlog, err := u.blogsRepo.Update(ctx, blog)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

// Delete blog
func (u *blogsUC) Delete(ctx context.Context, blogID uuid.UUID) error {

	if err := u.blogsRepo.Delete(ctx, blogID); err != nil {
		return err
	}

	return nil
}

// GetByID blog
func (u *blogsUC) GetByID(ctx context.Context, blogID uuid.UUID) (*models.Blog, error) {

	return u.blogsRepo.GetByID(ctx, blogID)
}

// GetAll blogs
func (u *blogsUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.BlogsList, error) {
	return u.blogsRepo.GetAll(ctx, title, query)
}
