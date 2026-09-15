package model

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model

	UserID    uint      `gorm:"not null"`
	CodeHash  string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`

	User User
}