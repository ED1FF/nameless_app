package main

import (
	"log"
	configs "nameless_app/configs"
	controllers "nameless_app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()
	configs.Connect()

	// routes------------------------------------------
	router.GET("", root)

	// user routes=====================================
	user := router.Group("user")
	{
		user.GET("", controllers.GetUser)
	}

	// run---------------------------------------------
	router.Run(":3000")
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
