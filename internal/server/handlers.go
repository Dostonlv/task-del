package server

import (
	"github.com/Dostonlv/task-del/docs"
	blogsHttp "github.com/Dostonlv/task-del/internal/blogs/delivery/http"
	newsHttp "github.com/Dostonlv/task-del/internal/news/delivery/http"

	"github.com/Dostonlv/task-del/internal/blogs/repository"
	"github.com/Dostonlv/task-del/internal/blogs/usecase"
	newRepo "github.com/Dostonlv/task-del/internal/news/repository"
	newUseCase "github.com/Dostonlv/task-del/internal/news/usecase"
	"github.com/Dostonlv/task-del/pkg/csrf"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Summary Health check endpoint
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} model "{"status": "Healthy!"}"
// @Router /health [get]
// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	bRepo := repository.NewBlogsRepository(s.db)
	nRepo := newRepo.NewNewsRepository(s.db)

	commUC := usecase.NewBlogsUseCase(s.cfg, bRepo, s.logger)
	newUC := newUseCase.NewNewsUseCase(nRepo, s.logger, s.cfg)

	// Init handlers
	blogHandlers := blogsHttp.NewBlogsHandlers(s.cfg, commUC, s.logger)
	newsHandlers := newsHttp.NewNewsHandlers(s.cfg, newUC, s.logger)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "blog and news API"
	docs.SwaggerInfo.Description = "blog and news REST API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/v1")

	health := v1.Group("/health")
	blogGroup := v1.Group("/blogs")
	newsGroup := v1.Group("/news")

	blogsHttp.MapBlogsRoutes(blogGroup, blogHandlers)
	newsHttp.MapNewsRoutes(newsGroup, newsHandlers)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy!"})
	})

	return nil
}
