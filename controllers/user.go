package user

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.JSON(200, gin.H{
		"message": name,
	})
}
