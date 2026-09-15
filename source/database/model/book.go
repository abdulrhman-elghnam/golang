package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title            string `gorm:"not null"`
	ISBN             string `gorm:"unique;not null"`
	Quantity         int    `gorm:"not null"`
	AvailableQuantity int    `gorm:"not null"`

	AuthorID   uint
	CategoryID uint

	Author   Author
	Category Category
}