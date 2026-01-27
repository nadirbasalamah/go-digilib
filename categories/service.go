package categories

import (
	"context"

	"gorm.io/gorm"
)

type Service interface {
	// GetAll(ctx context.Context) ([]Category, error)
	// GetByID(ctx context.Context, id uint) (Category, error)
	Create(ctx context.Context, category *CategoryRequest) (Category, error)
	// Update(ctx context.Context, category *CategoryRequest, id uint) (Category, error)
	// Delete(ctx context.Context, id uint) error
}

type service struct {
	create
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create: create{repository: repository},
	}
}
