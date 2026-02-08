package utils

import (
	"gorm.io/gorm"
)

// TODO: create DB scope for checking if the current user can update / delete the cart
func CurrentUser(userID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userID)
	}
}
