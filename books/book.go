package books

import (
	"go-digilib/categories"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Publisher   string              `json:"publisher"`
	Year        uint                `json:"year"`
	Stock       uint                `json:"stock"`
	CategoryID  uint                `json:"category_id"`
	Category    categories.Category `json:"category"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}

type BookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Publisher   string `json:"publisher" validate:"required"`
	Year        uint   `json:"year" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}
