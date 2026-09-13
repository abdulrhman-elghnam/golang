package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}