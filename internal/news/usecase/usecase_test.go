package usecase

import (
	"context"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/internal/news/mock"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/Dostonlv/task-del/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock

func TestNewUC_Create(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(mockNewRepo, logger, nil)

	// model of new
	new := models.New{}

	// context
	ctx := context.Background()

	// mock the Create method of the repository
	mockNewRepo.EXPECT().Create(
		ctx,
		gomock.Eq(&new),
	).Return(&new, nil)

	// call the Create method of the usecase
	createdNew, err := newUC.Create(context.Background(), &new)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, createdNew)
}

func TestNewUC_Update(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(mockNewRepo, logger, nil)

	// model of new
	new := models.New{
		ID:    uuid.New(),
		Title: "update-title",
	}

	// mock the Update method of the repository
	mockNewRepo.EXPECT().Update(
		context.Background(),
		gomock.Eq(&new),
	).Return(&new, nil)

	// call the Update method of the usecase
	updatedNew, err := newUC.Update(context.Background(), &new)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, updatedNew)
}

func TestNewUC_Delete(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(mockNewRepo, logger, nil)

	// new id
	newID := uuid.New()

	// mock the Delete method of the repository
	mockNewRepo.EXPECT().Delete(
		context.Background(),
		gomock.Eq(newID),
	).Return(nil)

	// call the Delete method of the usecase
	err := newUC.Delete(context.Background(), newID)

	// check the result
	require.NoError(t, err)
}

func TestNewUC_GetByID(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(mockNewRepo, logger, nil)

	// new id
	newID := uuid.New()

	// mock the GetByID method of the repository
	mockNewRepo.EXPECT().GetByID(
		context.Background(),
		gomock.Eq(newID),
	).Return(&models.New{}, nil)

	// call the GetByID method of the usecase
	new, err := newUC.GetByID(context.Background(), newID)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, new)
}

func TestNewUC_GetAll(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(mockNewRepo, logger, nil)

	// entity of NEW list, context, query
	entity := models.NewsList{}
	ctx := context.Background()
	query := utils.PaginationQuery{
		Page: 1,
		Size: 10,
	}

	// mock the GetAll method of the repository
	mockNewRepo.EXPECT().GetAll(
		ctx,
		&query,
	).Return(&entity, nil)

	// call the GetAll method of the usecase
	newList, err := newUC.GetAll(ctx, "", &query)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, newList)
}
