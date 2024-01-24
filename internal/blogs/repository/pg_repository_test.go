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

func TestBlogRepo_Create(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	uuid1 := uuid.New()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Create a blog success case
	t.Run("Create", func(t *testing.T) {
		// temprorary blog
		blog := &models.Blog{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blog.ID,
			blog.Title,
			blog.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(`INSERT INTO blogs (id,title,content) VALUES ($1,$2,$3) RETURNING *`).
			WithArgs(
				blog.Title,
				blog.Content,
			).WillReturnRows(rows)

		// call Create method
		createdBlog, err := repo.Create(context.Background(), blog)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, createdBlog)
		require.Equal(t, blog.ID, createdBlog.ID)
		require.Equal(t, blog.Title, createdBlog.Title)
		require.Equal(t, blog.Content, createdBlog.Content)
	})

	// Create blog error case
	t.Run("Create Error", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(
			`INSERT INTO blogs (id,title,content) VALUES ($1,$2,$3) RETURNING *;`,
		).WithArgs(
			blog.ID,
			blog.Title,
			blog.Content,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Create method
		createdBlog, err := repo.Create(context.Background(), blog)

		// check error and result
		require.Error(t, err)
		require.Nil(t, createdBlog)
	})
}

func TestBlogRepo_Update(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	uuid1 := uuid.New()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Update a blog success case
	t.Run("Update", func(t *testing.T) {
		// temprorary blog
		blog := &models.Blog{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blog.ID,
			blog.Title,
			blog.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`UPDATE blogs SET title = $1, content = $2 WHERE id = $3 RETURNING *`,
		).WithArgs(
			blog.Title,
			blog.Content,
			blog.ID,
		).WillReturnRows(rows)

		// call Update method
		updatedBlog, err := repo.Update(context.Background(), blog)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, updatedBlog)
		require.Equal(t, blog.ID, updatedBlog.ID)
		require.Equal(t, blog.Title, updatedBlog.Title)
		require.Equal(t, blog.Content, updatedBlog.Content)
	})

	// Update blog error case
	t.Run("Update Error", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      uuid1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(
			`UPDATE blogs SET title = $1, content = $2 WHERE id = $3 RETURNING *`,
		).WithArgs(
			blog.Title,
			blog.Content,
			blog.ID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Update method
		updatedBlog, err := repo.Update(context.Background(), blog)

		// check error and result
		require.Error(t, err)
		require.Nil(t, updatedBlog)
	})
}

// TestBlogRepo_Delete tests Delete method.
func TestBlogRepo_Delete(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Delete blog success case
	t.Run("Delete", func(t *testing.T) {

		// delete blog id
		blogID := uuid.New()

		// mock query with args and return result
		mock.ExpectExec(
			`DELETE FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.NoError(t, err)
	})

	// Delete blog error case
	t.Run("Delete Error", func(t *testing.T) {

		// delete blog id
		blogID := uuid.New()

		// mock query with args and return error
		mock.ExpectExec(
			`DELETE FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})

	// Delete blog RowsAffected equal to zero case
	t.Run("Delete RowsAffected equal to zero", func(t *testing.T) {

		// delete blog id
		blogID := uuid.New()

		// mock query with args and return result, but rows affected equal to zero
		mock.ExpectExec(
			`DELETE FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewResult(1, 0))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})

	// Delete blog RowsAffected error case
	t.Run("Delete RowsAffected Error", func(t *testing.T) {

		// delete blog id
		blogID := uuid.New()

		// mock query with args and return error which rows affected
		mock.ExpectExec(
			`DELETE FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})
}

// TestBlogRepo_GetByID tests GetByID method.
func TestBlogRepo_GetByID(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// GetByID success case
	t.Run("GetByID", func(t *testing.T) {

		// blog id
		blogID := uuid.New()

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blogID,
			"test-title",
			"test-content",
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`SELECT id, title, content, created_at FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnRows(rows)

		// call GetByID method
		blog, err := repo.GetByID(context.Background(), blogID)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blog)
		require.Equal(t, blogID, blog.ID)
		require.Equal(t, "test-title", blog.Title)
		require.Equal(t, "test-content", blog.Content)
	})

	// GetByID error case
	t.Run("GetByID Error", func(t *testing.T) {

		// blog id
		blogID := uuid.New()

		// mock query with args and return error
		mock.ExpectQuery(
			`SELECT id, title, content,created_at FROM blogs WHERE id = $1`,
		).WithArgs(
			blogID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetByID method
		blog, err := repo.GetByID(context.Background(), blogID)

		// check error and result
		require.Error(t, err)
		require.Nil(t, blog)
	})
}

// TestBlogRepo_GetAll tests GetAll method.
func TestBlogRepo_GetAll(t *testing.T) {

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// GetByID success case
	t.Run("GetAll", func(t *testing.T) {

		// blog id
		blogID := uuid.New()

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blogID,
			"test-title",
			"test-content",
		)

		// mock query with args and return rows
		mock.ExpectQuery(
			`SELECT id, title, content, created_at FROM blogs`,
		).WillReturnRows(rows)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), "", nil)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blogs)
		require.Equal(t, 1, len(blogs.Blogs))
		require.Equal(t, blogID, blogs.Blogs[0].ID)
		require.Equal(t, "test-title", blogs.Blogs[0].Title)
		require.Equal(t, "test-content", blogs.Blogs[0].Content)
	})

	// GetByID error case
	t.Run("GetAll Error", func(t *testing.T) {

		// mock query with args and return error
		mock.ExpectQuery(
			`SELECT id, title, content,created_at FROM blogs`,
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), "", nil)

		// check error and result
		require.Error(t, err)
		require.Nil(t, blogs)
	})

}
