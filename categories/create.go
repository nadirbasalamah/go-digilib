package categories

import (
	"context"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) Create(ctx context.Context, category *CategoryRequest) (Category, error) {
	result := c.repository.WithContext(ctx).Create(category)
	record := new(Category)

	if err := result.Error; err != nil {
		return Category{}, err
	}

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		return Category{}, err
	}

	return *record, nil
}
