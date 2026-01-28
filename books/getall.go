package books

import (
	"context"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAll(ctx context.Context) ([]Book, error) {
	books := []Book{}

	if err := g.repository.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
