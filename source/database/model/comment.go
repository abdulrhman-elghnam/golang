package model

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	PostID    uint      `gorm:"not null" json:"postId"`
	UserID    uint      `gorm:"not null" json:"userId"`
	Post      Post      `gorm:"foreignKey:PostID" json:"post"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}