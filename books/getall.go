package books

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAll(ctx context.Context) ([]Book, error) {
	books := []Book{}

	if err := g.repository.WithContext(ctx).Preload(clause.Associations).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
