package http

import (
	"github.com/Dostonlv/task-del/internal/blogs"
	"github.com/labstack/echo/v4"
)

// Map blogs routes
func MapBlogsRoutes(blogGroup *echo.Group, h blogs.Handlers) {
	blogGroup.POST("", h.Create())
	blogGroup.GET("", h.GetAll())
	blogGroup.DELETE("/:id", h.Delete())
	blogGroup.PUT("/:id", h.Update())
	blogGroup.GET("/:id", h.GetByID())
}
