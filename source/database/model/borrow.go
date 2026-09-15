package model

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model

	UserID uint
	BookID uint

	BorrowedAt time.Time
	DueDate    time.Time
	ReturnedAt *time.Time

	User User
	Book Book
}