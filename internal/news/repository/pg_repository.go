package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/internal/news"
	"github.com/Dostonlv/task-del/pkg/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// news Repository
type newsRepo struct {
	db *sqlx.DB
}

// NewNewsRepository News Repository constructor
func NewNewsRepository(db *sqlx.DB) news.Repository {
	return &newsRepo{db: db}
}

// Create new
func (r *newsRepo) Create(ctx context.Context, news *models.New) (*models.New, error) {
	newUUID := uuid.New()
	c := &models.New{}
	createNew := `INSERT INTO news (id,title,content) VALUES ($1,$2,$3) RETURNING *`
	if err := r.db.QueryRowxContext(
		ctx,
		createNew,
		newUUID,
		&news.Title,
		&news.Content,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Create.StructScan")
	}

	return c, nil
}

// Update new
func (r *newsRepo) Update(ctx context.Context, news *models.New) (*models.New, error) {
	updateNew := `UPDATE news SET
		title = $1,
		content = $2
	WHERE id = $3
	RETURNING *`
	res := &models.New{}
	if err := r.db.QueryRowxContext(ctx, updateNew, &news.Title, &news.Content, &news.ID).StructScan(res); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Update.QueryRowxContext")
	}

	return res, nil
}

// Delete new
func (r *newsRepo) Delete(ctx context.Context, ID uuid.UUID) error {
	deleteNew := `DELETE FROM news WHERE id = $1`

	result, err := r.db.ExecContext(ctx, deleteNew, ID)
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "newsRepo.Delete.rowsAffected")
	}

	return nil

}

// GetByID new
func (r *newsRepo) GetByID(ctx context.Context, newID uuid.UUID) (*models.New, error) {
	getNew := `SELECT id, title, content, created_at FROM news WHERE id = $1`
	new := &models.New{}
	if err := r.db.GetContext(ctx, new, getNew, newID); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetByID.GetContext")
	}

	return new, nil

}

// GetAll news
func (r *newsRepo) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error) {
	var (
		totalCount    int
		getTotalCount = `SELECT COUNT(id) FROM news`
		getAllNews    = `SELECT id, title, content ,created_at
							FROM news`
	)

	if title != "" {
		getTotalCount = fmt.Sprintf("%s%s", getTotalCount, " and title LIKE '%"+title+"%';")
		getAllNews = fmt.Sprintf("%s%s", getAllNews, " and title LIKE '%"+title+"%' ")
	}

	getAllNews += " ORDER BY created_at OFFSET $1 LIMIT $2;"
	if err := r.db.QueryRowContext(ctx, getTotalCount).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryRowContext")
	}

	if totalCount == 0 {
		return &models.NewsList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			News:       make([]*models.New, 0),
		}, nil

	}
	rows, err := r.db.QueryxContext(ctx, getAllNews, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	newsList := make([]*models.New, 0, query.GetSize())
	for rows.Next() {
		new := &models.New{}
		if err := rows.StructScan(new); err != nil {
			return nil, errors.Wrap(err, "newsRepo.GetAll.StructScan")
		}
		newsList = append(newsList, new)
	}

	return &models.NewsList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		News:       newsList,
	}, nil
}
