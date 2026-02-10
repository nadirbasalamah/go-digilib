package rents

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getbyuser struct {
	repository *gorm.DB
}

func (g getbyuser) GetByUser(ctx context.Context, userId uint) ([]Rent, error) {
	rents := []Rent{}

	if err := g.repository.
		WithContext(ctx).
		Preload(clause.Associations).
		Where("user_id = ?", userId).
		Find(&rents).
		Error; err != nil {
		return nil, err
	}

	return rents, nil
}
