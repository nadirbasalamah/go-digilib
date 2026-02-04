package rents

import (
	"go-digilib/books"
	"go-digilib/users"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	ID         uint           `gorm:"primaryKey"`
	BookID     uint           `json:"book_id"`
	Book       books.Book     `json:"book"`
	UserID     uint           `json:"user_id"`
	User       users.User     `json:"user"`
	Quantity   uint           `json:"quantity"`
	Fee        float64        `json:"fee"`
	Status     string         `json:"rent_status" gorm:"type:rent_status"`
	ReturnTime time.Time      `json:"return_time"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	ReturnedAt time.Time      `json:"returned_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RentRequest struct {
	BookID   uint `json:"book_id" validate:"required"`
	UserID   uint
	Quantity uint `json:"quantity" validate:"required,gte=1"`
}

type RentsRequest struct {
	Rents []RentRequest
}
