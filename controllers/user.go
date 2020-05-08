package controllers

import (
	"nameless_app/model"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.JSON(200, gin.H{
		"message": name,
	})
}

func PostUser(c *gin.Context) {
	c.Request.ParseMultipartForm(10)
	user, err := model.CreateUser(c.Request.PostForm)

	if err != "" {
		c.JSON(503, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{
			"message": "User Successfully Created",
			"record":  user,
		})
	}
}

func GetUsers(c *gin.Context) {
	users := model.GetAllUsers()
	c.JSON(200, users)
}
