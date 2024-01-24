package http

import (
	"github.com/Dostonlv/task-del/config"
	"github.com/Dostonlv/task-del/internal/blogs"
	"github.com/Dostonlv/task-del/internal/models"
	"github.com/Dostonlv/task-del/pkg/httpErrors"
	"github.com/Dostonlv/task-del/pkg/logger"
	"github.com/Dostonlv/task-del/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// blogs handlers
type blogsHandlers struct {
	cfg    *config.Config
	blogUC blogs.UseCase
	logger logger.Logger
}

// NewBlogsHandlers Blogs handlers constructor
func NewBlogsHandlers(cfg *config.Config, blogsUC blogs.UseCase, logger logger.Logger) blogs.Handlers {
	return &blogsHandlers{cfg: cfg, blogUC: blogsUC, logger: logger}
}

// Create
// @Summary Create new blog
// @Description create new blog
// @Tags blogs
// @Accept json
// @Produce json
// @Param body body models.BlogsSwagger true "body"
// @Success 201 {object} models.Blog
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs [post]
func (h *blogsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		blog := &models.Blog{}

		if err := utils.SanitizeRequest(c, blog); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		createdblog, err := h.blogUC.Create(c.Request().Context(), blog)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdblog)
	}
}

// Update
// @Summary Update blog
// @Description update new blog
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.BlogsSwagger true "body"
// @Success 200 {object} models.BlogsSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [put]
func (h *blogsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &models.Blog{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedblog, err := h.blogUC.Update(c.Request().Context(), &models.Blog{
			ID:      blogsID,
			Title:   comm.Title,
			Content: comm.Content,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedblog)
	}
}

// Delete
// @Summary Delete blog
// @Description delete blog
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [delete]
func (h *blogsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.blogUC.Delete(c.Request().Context(), blogsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get blog
// @Description Get blog by id
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Blog
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs/{id} [get]
func (h *blogsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		blog, err := h.blogUC.GetByID(c.Request().Context(), blogsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, blog)
	}
}

// GetAll
// @Summary Get blogs
// @Description Get all blog
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param title query string false "title"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Success 200 {object} models.BlogsList
// @Failure 500 {object} httpErrors.RestErr
// @Router /blogs [get]
func (h *blogsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		blogList, err := h.blogUC.GetAll(c.Request().Context(), c.QueryParam("title"), pq)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, blogList)
	}
}
