package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/abdulrhman-elghnam/golang/source/database"
	"github.com/abdulrhman-elghnam/golang/source/module/authentication"
)
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	 err := database.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	app := router.Group("/")

	authentication.AuthenticationControllerRegistration(app)

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from backend server 🚀",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route is not exist 🦦",
		})
	})
	router.Run(":" + os.Getenv("PORT"))
}
