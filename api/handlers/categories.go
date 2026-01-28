package handlers

import (
	"go-digilib/categories"
	"go-digilib/shared/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Categories struct {
	categories categories.Service
}

func (c Categories) Create(ctx *echo.Context) error {
	categoryReq := new(categories.CategoryRequest)

	if err := ctx.Bind(categoryReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(categoryReq); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, models.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	category, err := c.categories.Create(ctx.Request().Context(), categoryReq)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response[any]{
			Status:  "failed",
			Message: "create category failed",
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response[categories.Category]{
		Status:  "success",
		Message: "category created",
		Data:    category,
	})
}

func NewCategories(categories categories.Service) Categories {
	categoriesHandler := Categories{
		categories: categories,
	}

	return categoriesHandler
}
