package api

import (
	"go-digilib/api/handlers"
	"go-digilib/api/middlewares"
	"go-digilib/categories"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB) *echo.Echo {
	var (
		e                 = echo.New()
		categories        = categories.New(repository)
		categoriesHandler = handlers.NewCategories(categories)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	categoryRoutes := e.Group("/api/v1")

	categoryRoutes.POST("/categories", categoriesHandler.Create)

	return e
}
