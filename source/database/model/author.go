package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	Books []Book
}