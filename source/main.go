package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/abdulrhman-elghnam/golang/source/configuration"
	"github.com/abdulrhman-elghnam/golang/source/database"
	"github.com/abdulrhman-elghnam/golang/source/module/authentication"
)

func main() {
	err := configuration.LoadConfig()
	if err != nil {
		log.Fatal("error while using .env")
	}

	

	db, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	app := router.Group("/")

	authentication.AuthenticationControllerRegistration(app,db)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello from backend server 🚀",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "route is not exist 🦦",
		})
	})
	router.Run(":" + os.Getenv("PORT"))
}
