package handlers

import (
	"go-digilib/books"
	"go-digilib/shared/dtos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Books struct {
	books books.Service
}

func (b Books) GetAll(ctx *echo.Context) error {
	booksData, err := b.books.GetAll(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch books failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]books.Book]{
		Status:  "success",
		Message: "all books",
		Data:    booksData,
	})
}

func (b Books) GetByID(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	book, err := b.books.GetByID(ctx.Request().Context(), uint(id))

	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[books.Book]{
		Status:  "success",
		Message: "book found",
		Data:    book,
	})
}

func (b Books) Create(ctx *echo.Context) error {
	bookReq := new(books.BookRequest)

	if err := ctx.Bind(bookReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(bookReq); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	book, err := b.books.Create(ctx.Request().Context(), bookReq)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "create book failed",
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[books.Book]{
		Status:  "success",
		Message: "book created",
		Data:    book,
	})
}

func (b Books) Update(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	bookReq := new(books.BookRequest)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	if err := ctx.Bind(bookReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(bookReq); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	book, err := b.books.Update(ctx.Request().Context(), bookReq, uint(id))

	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "update book failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[books.Book]{
		Status:  "success",
		Message: "book updated",
		Data:    book,
	})
}

func (b Books) Delete(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	err = b.books.Delete(ctx.Request().Context(), uint(id))

	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "success",
		Message: "book deleted",
	})
}

func NewBooks(books books.Service) Books {
	booksHandler := Books{
		books: books,
	}

	return booksHandler
}
