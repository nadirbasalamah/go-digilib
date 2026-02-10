package rents

import (
	"context"
	"go-digilib/db/models"
	"time"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	get
}

func (u update) Update(ctx context.Context, rentReq *RentUpdateRequest, id uint) (Rent, error) {
	status := models.RentStatus(rentReq.Status)

	rent := models.Rent{
		Status: status,
	}

	if status == models.Returned {
		rent.ReturnedAt = time.Now()
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&rent)

	if err := result.Error; err != nil {
		return Rent{}, err
	}

	record, err := u.get.GetByID(ctx, id)

	if err != nil {
		return Rent{}, err
	}

	return record, nil
}
