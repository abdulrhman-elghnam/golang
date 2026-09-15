package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Phone     string    `gorm:"unique;not null"`
	DOB       time.Time `gorm:"not null"`
	Role      string    `gorm:"not null;default:member"`

	Borrows []Borrow
}