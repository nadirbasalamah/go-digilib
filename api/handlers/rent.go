package handlers

import (
	"go-digilib/api/middlewares"
	"go-digilib/pkg/dtos"
	"go-digilib/rents"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Rents struct {
	rents rents.Service
}

func (r Rents) GetByUser(ctx *echo.Context) error {
	userData := ctx.Get("userData").(*middlewares.JWTCustomClaims)
	userID := userData.ID

	rentsData, err := r.rents.GetByUser(ctx.Request().Context(), uint(userID))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch rents failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]rents.Rent]{
		Status:  "success",
		Message: "all rents",
		Data:    rentsData,
	})
}

func (r Rents) Create(ctx *echo.Context) error {
	userData := ctx.Get("userData").(*middlewares.JWTCustomClaims)
	userID := userData.ID

	rentReq := ctx.Get("validatedBody").(*rents.RentRequest)
	rentReq.UserID = uint(userID)

	//TODO: calculate ongkir from RajaOngkir API

	//TODO: calculate return time in timestamp format

	rent, err := r.rents.Create(ctx.Request().Context(), rentReq)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "add book to the rent failed",
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[rents.Rent]{
		Status:  "success",
		Message: "book added to the rent",
		Data:    rent,
	})
}

func (r Rents) Update(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	rentReq := ctx.Get("validatedBody").(*rents.RentUpdateRequest)

	rent, err := r.rents.Update(ctx.Request().Context(), rentReq, uint(id))

	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "update rent failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[rents.Rent]{
		Status:  "success",
		Message: "rent updated",
		Data:    rent,
	})
}

func (r Rents) Delete(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	err = r.rents.Delete(ctx.Request().Context(), uint(id))

	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "cart not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "success",
		Message: "book removed from the cart",
	})
}

func NewRents(rents rents.Service) Rents {
	rentsHandler := Rents{
		rents: rents,
	}

	return rentsHandler
}
