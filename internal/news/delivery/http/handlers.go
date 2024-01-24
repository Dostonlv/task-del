package http

import (
	"github.com/Dostonlv/task-del/config"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/internal/news"
	"github.com/Dostonlv/task-del/pkg/httpErrors"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/Dostonlv/task-del/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

// news handlers
type newsHandlers struct {
	cfg    *config.Config
	newsUC news.UseCase
	logger logger.Logger
}

// NewNewsHandlers News handlers constructor
func NewNewsHandlers(cfg *config.Config, newsUC news.UseCase, logger logger.Logger) news.Handlers {
	return &newsHandlers{cfg: cfg, newsUC: newsUC, logger: logger}
}

// Create
// @Summary Create new news
// @Description create new news
// @Tags news
// @Accept json
// @Produce json
// @Param body body models.NewsSwagger true "body"
// @Success 201 {object} models.New
// @Failure 500 {object} httpErrors.RestErr
// @Router /news [post]
func (h *newsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		news := &models.New{}

		if err := utils.SanitizeRequest(c, news); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		creatednews, err := h.newsUC.Create(c.Request().Context(), news)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, creatednews)
	}
}

// Update
// @Summary Update news
// @Description update new news
// @Tags news
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param body body models.NewsSwagger true "body"
// @Success 200 {object} models.NewsSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /news/{id} [put]
func (h *newsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		newID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &models.New{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatednews, err := h.newsUC.Update(c.Request().Context(), &models.New{
			ID:      newID,
			Title:   comm.Title,
			Content: comm.Content,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatednews)
	}
}

// Delete
// @Summary Delete news
// @Description deleted news
// @Tags news
// @Accept json
// @Produce json
// @Param id path string true "news ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /news/{id} [delete]
func (h *newsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err := h.newsUC.Delete(c.Request().Context(), newsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get news by ID
// @Description get news by ID
// @Tags news
// @Accept json
// @Produce json
// @Param id path string true "news ID"
// @Success 200 {object} models.New
// @Failure 500 {object} string
// @Router /news/{id} [get]
func (h *newsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		newsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		news, err := h.newsUC.GetByID(c.Request().Context(), newsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, news)
	}

}

// GetAll
// @Summary Get all news
// @Description get all news
// @Tags news
// @Accept json
// @Produce json
// @Param title query string false "title"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} models.NewsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /news [get]
func (h *newsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		newList, err := h.newsUC.GetAll(c.Request().Context(), c.QueryParam("title"), pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, newList)
	}
}
