package carts

import (
	"context"
	"go-digilib/db/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	get
}

func (u update) Update(ctx context.Context, cartReq *CartRequest, id uint) (Cart, error) {
	cart := models.Cart{
		BookID:   cartReq.BookID,
		Quantity: cartReq.Quantity,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&cart)

	if err := result.Error; err != nil {
		return Cart{}, nil
	}

	record, err := u.get.GetByID(ctx, id)

	if err != nil {
		return Cart{}, err
	}

	return record, nil
}
