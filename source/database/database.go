package database

import (
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(".app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
