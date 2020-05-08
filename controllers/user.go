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
	user, messege := model.CreateUser(c.Request.PostForm)
	c.JSON(200, gin.H{
		"message": messege,
		"record":  user,
	})
}

func GetUsers(c *gin.Context) {
	users := model.GetAllUsers()
	c.JSON(200, users)
}
