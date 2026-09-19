package authentication

import (
	"errors"
	"net/http"

	jwtpkg "github.com/abdulrhman-elghnam/golang/source/common/security/jwt"
	security "github.com/abdulrhman-elghnam/golang/source/common/security/private"
	"github.com/abdulrhman-elghnam/golang/source/common/structure"
	"github.com/abdulrhman-elghnam/golang/source/database/model"
	"github.com/abdulrhman-elghnam/golang/source/database/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(userRepo *repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		var request SignUpDTO

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

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Database error",
			)
			return
		}

		hashedPassword, err := security.HashPassword(request.Password)
		if err != nil {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Failed to hash password",
			)
			return
		}

		user = model.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  hashedPassword,
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

func LogIn(userRepo *repository.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request LogInDTO
		var user model.User

		if err := ctx.ShouldBindJSON(&request); err != nil {
			structure.Fail(
				ctx,
				http.StatusBadRequest,
				err.Error(),
			)
			return
		}

		err := userRepo.FindOne(
			map[string]any{
				"email": request.Email,
			},
			&user,
		)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			structure.Fail(
				ctx,
				http.StatusUnauthorized,
				"Invalid email or password",
			)
			return
		}

		if err != nil {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Database error",
			)
			return
		}

		if !security.ComparePassword(
			request.Password,
			user.Password,
		) {
			structure.Fail(
				ctx,
				http.StatusUnauthorized,
				"Invalid email or password",
			)
			return
		}

		token, err := jwtpkg.GenerateToken(user.ID)
		if err != nil {
			structure.Fail(
				ctx,
				http.StatusInternalServerError,
				"Failed to generate token",
			)
			return
		}

		structure.OK(ctx, gin.H{
			"message": "Login successful",
			"user": gin.H{
				"id":        user.ID,
				"firstName": user.FirstName,
				"lastName":  user.LastName,
				"email":     user.Email,
				"phone":     user.Phone,
				"role":      user.Role,
			},
			"token": token,
		}, http.StatusOK)
	}
}

func ConfirmEmail(userRepo *repository.Repository) {

}

func ForgetPassword(userRepo *repository.Repository) {

}

func ChangePassword(userRepo *repository.Repository) {

}

func ConfirmPassword(userRepo *repository.Repository) {

}
