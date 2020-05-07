package main

import (
	user "nameless_app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", root)
	router.GET("/user", user.GetUser)

	router.Run(":3000")
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
