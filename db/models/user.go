package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `json:"username" gorm:"unique"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"-"`
	Address        string `json:"address" gorm:"type:text"`
	ProfilePicture string `json:"profile_picture"`
	//TODO: add role field
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
