package authentication

import "github.com/gin-gonic/gin"

func AuthenticationControllerRegistration(rg *gin.RouterGroup) {

	router := rg.Group("/authentication")

	router.GET("/signup", SignUp())

}
