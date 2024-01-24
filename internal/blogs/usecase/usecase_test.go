package usecase

import (
	"context"
	"github.com/Dostonlv/task-del/internal/blogs/mock"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

// use gomock to generate mock for usecase
//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock

func TestBlofUC_Create(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockUseCase(ctrl)
	blogUC := NewBlogsUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{}

	// context
	ctx := context.Background()

	// mock the Create method of the repository
	mockBlogRepo.EXPECT().Create(
		ctx,
		gomock.Eq(&blog),
	).Return(&blog, nil)

	// call the Create method of the usecase
	createdBlog, err := blogUC.Create(context.Background(), &blog)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, createdBlog)
}

func TestBlofUC_Update(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockUseCase(ctrl)
	blogUC := NewBlogsUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{}

	// context
	ctx := context.Background()

	// mock the Update method of the repository
	mockBlogRepo.EXPECT().Update(
		ctx,
		gomock.Eq(&blog),
	).Return(&blog, nil)

	// call the Update method of the usecase
	updatedBlog, err := blogUC.Update(context.Background(), &blog)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, updatedBlog)
}

func TestBlofUC_Delete(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockUseCase(ctrl)
	blogUC := NewBlogsUseCase(nil, mockBlogRepo, logger)

	// context
	ctx := context.Background()

	// mock the Delete method of the repository
	mockBlogRepo.EXPECT().Delete(
		ctx,
		gomock.Any(),
	).Return(nil)

	// call the Delete method of the usecase
	err := blogUC.Delete(context.Background(), uuid.New())

	// check the result
	require.NoError(t, err)
}

func TestBlofUC_GetByID(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockUseCase(ctrl)
	blogUC := NewBlogsUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{}

	// context
	ctx := context.Background()

	// mock the GetByID method of the repository
	mockBlogRepo.EXPECT().GetByID(
		ctx,
		gomock.Any(),
	).Return(&blog, nil)

	// call the GetByID method of the usecase
	blog1, err := blogUC.GetByID(context.Background(), uuid.New())

	// check the result
	require.NoError(t, err)
	require.NotNil(t, blog1)
}

func TestBlofUC_GetAll(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockUseCase(ctrl)
	blogUC := NewBlogsUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{}

	// context
	ctx := context.Background()

	// mock the GetAll method of the repository
	mockBlogRepo.EXPECT().GetAll(
		ctx,
		gomock.Any(),
		gomock.Any(),
	).Return(&models.BlogsList{Blogs: []*models.Blog{&blog}}, nil)

	// call the GetAll method of the usecase
	blogsList, err := blogUC.GetAll(context.Background(), "", nil)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, blogsList)
}
