package rents

import (
	"context"

	"gorm.io/gorm"
)

type delete struct {
	repository *gorm.DB
	get
}

func (d delete) Delete(ctx context.Context, id uint) error {
	rent, err := d.get.GetByID(ctx, id)

	if err != nil {
		return err
	}

	if err := d.repository.Delete(&rent).Error; err != nil {
		return err
	}

	return nil
}
