package books

import (
	"context"

	"gorm.io/gorm"
)

type get struct {
	repository *gorm.DB
}

func (g get) GetByID(ctx context.Context, id uint) (Book, error) {
	book := new(Book)

	if err := g.repository.WithContext(ctx).First(book, "id = ?", id).Error; err != nil {
		return Book{}, err
	}

	return *book, nil
}
