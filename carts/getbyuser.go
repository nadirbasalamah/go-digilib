package carts

import (
	"context"
	"go-digilib/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getbyuser struct {
	repository *gorm.DB
}

func (g getbyuser) GetByUser(ctx context.Context, pagination utils.Pagination, userId uint) (utils.Pagination, error) {
	carts := []Cart{}

	if err := g.repository.
		WithContext(ctx).
		Scopes(utils.PaginateByUserID(&carts, &pagination, userId, g.repository)).
		Preload(clause.Associations).
		Find(&carts).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = carts

	return pagination, nil
}
