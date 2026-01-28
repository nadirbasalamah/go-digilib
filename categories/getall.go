package categories

import (
	"context"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAll(ctx context.Context) ([]Category, error) {
	categories := []Category{}

	if err := g.repository.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
