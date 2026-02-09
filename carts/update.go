package carts

import (
	"context"
	"errors"
	"go-digilib/db/models"
	"go-digilib/pkg/utils"

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

	result := u.repository.WithContext(ctx).Scopes(utils.CurrentUser(cartReq.UserID)).Where("id = ?", id).Updates(&cart)

	isFailed := result.Error != nil || result.RowsAffected == 0

	if isFailed {
		return Cart{}, errors.New("update failed")
	}

	record, err := u.get.GetByID(ctx, id)

	if err != nil {
		return Cart{}, err
	}

	return record, nil
}
