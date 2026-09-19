package database

import (
	"errors"

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

	DB = db

	err = db.AutoMigrate(
		&model.User{},
		&model.Book{},
		&model.Author{},
		&model.Category{},
		&model.Borrow{},
		&model.OTP{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}


func FindUserByID(id uint) (*model.User, error) {
	if DB == nil {
		var err error
		DB, err = DatabaseConnection()
		if err != nil {
			return nil, err
		}
	}

	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
