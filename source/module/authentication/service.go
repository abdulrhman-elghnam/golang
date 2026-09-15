package authentication

import (
	"net/http"

	"github.com/abdulrhman-elghnam/golang/source/common/structure"
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"github.com/abdulrhman-elghnam/golang/source/database/repository"
	"github.com/abdulrhman-elghnam/golang/source/module/authentication/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SignUp(userRepo *repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		var request dto.SignUpDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			structure.Fail(ctx, http.StatusBadRequest, err.Error())
			return
		}

		err := userRepo.FindOne(
			map[string]any{
				"email": request.Email,
			},
			&user,
		)

		if err == nil {
			structure.Fail(
				ctx,
				http.StatusConflict,
				"Email already exists",
			)
			return
		}

		if err != gorm.ErrRecordNotFound {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Database error",
			)
			return
		}

		user = model.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  request.Password,
			Phone:     request.Phone,
		}

		if err := userRepo.Create(&user); err != nil {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Failed to create user",
			)
			return
		}

		structure.OK(ctx, gin.H{
			"message": "User created successfully",
		}, http.StatusCreated)
	}
}