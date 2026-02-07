package carts

import (
	"context"
	"go-digilib/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type create struct {
	repository *gorm.DB
}

func (c create) Create(ctx context.Context, cartReq *CartRequest) (Cart, error) {
	cart := models.Cart{
		BookID:   cartReq.BookID,
		UserID:   cartReq.UserID,
		Quantity: cartReq.Quantity,
	}

	result := c.repository.WithContext(ctx).Create(&cart)
	record := new(Cart)

	if err := result.Error; err != nil {
		return Cart{}, err
	}

	if err := result.WithContext(ctx).Preload(clause.Associations).Last(record).Error; err != nil {
		return Cart{}, err
	}

	return *record, nil
}
