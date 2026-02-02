package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Username       string `json:"username" gorm:"unique"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"-"`
	Address        string `json:"address" gorm:"type:text"`
	ProfilePicture string `json:"profile_picture"`
	// Role           models.Role    `json:"role" gorm:"type:role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsNumber,containsSpecialCharacter"`
	Address  string `json:"address" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
