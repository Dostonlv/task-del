package usecase

import (
	"context"
	"github.com/Dostonlv/task-del/config"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/internal/news"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/Dostonlv/task-del/pkg/utils"
	"github.com/google/uuid"
)

// news use case
type newsUC struct {
	newsRepo news.Repository
	logger   logger.Logger
	cfg      *config.Config
}

// NewNewsUseCase news use case constructor
func NewNewsUseCase(newsRepo news.Repository, logger logger.Logger, cfg *config.Config) news.UseCase {
	return &newsUC{newsRepo: newsRepo, logger: logger, cfg: cfg}
}

// Create news
func (u *newsUC) Create(ctx context.Context, news *models.New) (*models.New, error) {
	return u.newsRepo.Create(ctx, news)
}

// Update news
func (u *newsUC) Update(ctx context.Context, news *models.New) (*models.New, error) {
	updatedNews, err := u.newsRepo.Update(ctx, news)
	if err != nil {
		return nil, err
	}

	return updatedNews, nil
}

// Delete news
func (u *newsUC) Delete(ctx context.Context, newsID uuid.UUID) error {

	if err := u.newsRepo.Delete(ctx, newsID); err != nil {
		return err
	}

	return nil
}

// GetByID news
func (u *newsUC) GetByID(ctx context.Context, newsID uuid.UUID) (*models.New, error) {

	return u.newsRepo.GetByID(ctx, newsID)
}

// GetAll news
func (u *newsUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error) {
	return u.newsRepo.GetAll(ctx, title, query)
}
