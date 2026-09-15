package database

import (
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(".app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Book{},
		&model.Author{},
		&model.Category{},
		&model.Borrow{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
