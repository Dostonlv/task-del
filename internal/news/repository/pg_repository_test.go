package repository

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
)

//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock

// TestNewRepo_Create tests Create method.
func TestNewRepo_Create(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new's repository
	repo := NewNewsRepository(sqlxDB)

	// Create new success case
	t.Run("Create", func(t *testing.T) {
		uuid1 := uuid.New()
		// temprorary new
		new := &models.New{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			new.ID,
			new.Title,
			new.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`INSERT INTO news (id,title,content) VALUES ($1,$2,$3) RETURNING *`,
		).WithArgs(
			new.Title,
			new.Content,
		).WillReturnRows(rows)

		// call Create method
		createdNew, err := repo.Create(context.Background(), new)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, createdNew)
		require.Equal(t, new.ID, createdNew.ID)
		require.Equal(t, new.Title, createdNew.Title)
		require.Equal(t, new.Content, createdNew.Content)
	})

	// Create New error case
	t.Run("Create Error", func(t *testing.T) {

		// temprorary new
		new := &models.New{
			ID:      uuid.New(),
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(
			`INSERT INTO news (id,title,content) VALUES ($1,$2,$3) RETURNING *`,
		).WithArgs(
			new.ID,
			new.Title,
			new.Content,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Create method
		createdNew, err := repo.Create(context.Background(), new)

		// check error and result
		require.Error(t, err)
		require.Nil(t, createdNew)
	})
}

func TestNewRepo_Update(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// Update new success case
	t.Run("Update", func(t *testing.T) {

		// temprorary new
		new := &models.New{
			ID:      uuid.New(),
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			new.ID,
			new.Title,
			new.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`UPDATE news SET title = $1, content = $2 WHERE id = $3 RETURNING *`,
		).WithArgs(
			new.Title,
			new.Content,
			new.ID,
		).WillReturnRows(rows)

		// call Update method
		updatedNew, err := repo.Update(context.Background(), new)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, updatedNew)
		require.Equal(t, new.ID, updatedNew.ID)
		require.Equal(t, new.Title, updatedNew.Title)
		require.Equal(t, new.Content, updatedNew.Content)
	})

	// Update New error case
	t.Run("Update Error", func(t *testing.T) {
		uuid1 := uuid.New()
		// temprorary New
		new := &models.New{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(
			`UPDATE news SET title = $1, content = $2 WHERE id = $3 RETURNING *`,
		).WithArgs(
			new.Title,
			new.Content,
			new.ID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Update method
		updatedNew, err := repo.Update(context.Background(), new)

		// check error and result
		require.Error(t, err)
		require.Nil(t, updatedNew)
	})
}

// TestNewRepo_Delete tests Delete method.
func TestNewRepo_Delete(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// Delete newNew success case
	t.Run("Delete", func(t *testing.T) {

		// delete new id
		newID := uuid.New()

		// mock query with args and return result
		mock.ExpectExec(
			`DELETE FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.NoError(t, err)
	})

	// Delete new error case
	t.Run("Delete Error", func(t *testing.T) {

		// delete new id
		newID := uuid.New()

		// mock query with args and return error
		mock.ExpectExec(
			`DELETE FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})

	// Delete new RowsAffected equal to zero case
	t.Run("Delete RowsAffected equal to zero", func(t *testing.T) {

		// delete new id
		newID := uuid.New()

		// mock query with args and return result, but rows affected equal to zero
		mock.ExpectExec(
			`DELETE FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewResult(1, 0))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})

	// Delete new RowsAffected error case
	t.Run("Delete RowsAffected Error", func(t *testing.T) {

		// delete new id
		newID := uuid.New()

		// mock query with args and return error which rows affected
		mock.ExpectExec(
			`DELETE FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})
}

// TestNewRepo_GetByID tests GetByID method.
func TestNewRepo_GetByID(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// GetByID success case
	t.Run("GetByID", func(t *testing.T) {

		// new id
		newID := uuid.New()

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			newID,
			"test-title",
			"test-content",
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`SELECT id, title, content, created_at FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnRows(rows)

		// call GetByID method
		new, err := repo.GetByID(context.Background(), newID)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, new)
		require.Equal(t, newID, new.ID)
		require.Equal(t, "test-title", new.Title)
		require.Equal(t, "test-content", new.Content)
	})

	// GetByID error case
	t.Run("GetByID Error", func(t *testing.T) {

		// new id
		newID := uuid.New()

		// mock query with args and return error
		mock.ExpectQuery(
			`SELECT id, title, content, created_at FROM news WHERE id = $1`,
		).WithArgs(
			newID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetByID method
		new, err := repo.GetByID(context.Background(), newID)

		// check error and result
		require.Error(t, err)
		require.Nil(t, new)
	})
}

// TestNewRepo_GetAll tests GetAll method.
func TestNewRepo_GetAll(t *testing.T) {}
