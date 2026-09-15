package authentication

import (
	"errors"
	"net/http"

	"github.com/abdulrhman-elghnam/golang/source/common/security"
	"github.com/abdulrhman-elghnam/golang/source/common/structure"
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"github.com/abdulrhman-elghnam/golang/source/database/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp() gin.HandlerFunc {
	repo := repository.NewRepository()

return func(ctx *gin.Context) {

	var existingUser model.User

	user := model.User{
		Name:     "Abdulrhman",
		Email:    "abdulrhman@gmail.com",
		Role:     "user",
	}

	err := repo.FindOne(
		map[string]any{
			"email": user.Email,
		},
		&existingUser,
	)

	if err == nil {
		structure.Fail(
			ctx,
			http.StatusConflict,
			"Email already exists",
		)
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		structure.Fail(
			ctx,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	hashedPassword, err := security.HashPassword("1234")
	if err != nil {
		structure.Fail(
			ctx,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	user.Password = hashedPassword

	err = repo.Create(&user)
	if err != nil {
		structure.Fail(
			ctx,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
}
