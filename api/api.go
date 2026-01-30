package api

import (
	"go-digilib/api/handlers"
	"go-digilib/api/middlewares"
	"go-digilib/books"
	"go-digilib/categories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary) *echo.Echo {
	var (
		e                 = echo.New()
		categories        = categories.New(repository)
		books             = books.New(repository)
		categoriesHandler = handlers.NewCategories(categories)
		booksHandler      = handlers.NewBooks(books, cld)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	logger, _ := zap.NewProduction()
	loggerConfig := middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}

	logMiddleware := middlewares.LoggerConfig{Config: loggerConfig}

	e.Use(logMiddleware.Init())

	categoryRoutes := e.Group("/api/v1")

	categoryRoutes.GET("/categories", categoriesHandler.GetAll)
	categoryRoutes.GET("/categories/:id", categoriesHandler.GetByID)
	categoryRoutes.POST("/categories", categoriesHandler.Create)
	categoryRoutes.PATCH("/categories/:id", categoriesHandler.Update)
	categoryRoutes.DELETE("/categories/:id", categoriesHandler.Delete)

	bookRoutes := e.Group("/api/v1")
	bookRoutes.GET("/books", booksHandler.GetAll)
	bookRoutes.GET("/books/:id", booksHandler.GetByID)
	bookRoutes.POST("/books", booksHandler.Create)
	bookRoutes.PATCH("/books/:id", booksHandler.Update)
	bookRoutes.DELETE("/books/:id", booksHandler.Delete)

	return e
}
