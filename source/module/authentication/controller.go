package authentication

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthenticationControllerRegistration(rg *gin.RouterGroup, db *gorm.DB) {

	router := rg.Group("/authentication")

	router.GET("/signup", SignUp(db))

}
