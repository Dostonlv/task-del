package http

import (
	"github.com/Dostonlv/task-del/internal/news"
	"github.com/labstack/echo/v4"
)

// Map news routes
func MapNewsRoutes(newsGroup *echo.Group, h news.Handlers) {
	newsGroup.POST("", h.Create())
	newsGroup.GET("", h.GetAll())
	newsGroup.DELETE("/:id", h.Delete())
	newsGroup.PUT("/:id", h.Update())
	newsGroup.GET("/:id", h.GetByID())
}
