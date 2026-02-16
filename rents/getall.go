package rents

import (
	"context"
	"go-digilib/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAll(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	rents := []Rent{}

	if err := g.repository.
		WithContext(ctx).
		Scopes(utils.Paginate(&rents, &pagination, g.repository)).
		Preload(clause.Associations).
		Find(&rents).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = rents

	return pagination, nil
}
