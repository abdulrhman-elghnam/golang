package authentication

import (
	"github.com/abdulrhman-elghnam/golang/source/database/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthenticationControllerRegistration(
	rg *gin.RouterGroup,
	db *gorm.DB,
) {
	router := rg.Group("/authentication")

	userRepo := repository.NewRepository(db)

	router.POST("/signup", SignUp(userRepo))
}