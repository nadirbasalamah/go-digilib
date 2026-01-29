package api

import (
	"go-digilib/api/handlers"
	"go-digilib/api/middlewares"
	"go-digilib/books"
	"go-digilib/categories"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB) *echo.Echo {
	var (
		e                 = echo.New()
		categories        = categories.New(repository)
		books             = books.New(repository)
		categoriesHandler = handlers.NewCategories(categories)
		booksHandler      = handlers.NewBooks(books)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

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
