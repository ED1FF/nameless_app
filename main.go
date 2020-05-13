package main

import (
	"log"
	"nameless_app/configs"
	"nameless_app/controllers"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Thanks to otraore for the code example
// https://gist.github.com/otraore/4b3120aa70e1c1aa33ba78e886bb54f3

func main() {
	conf := configs.New()
	configs.Connect(conf.Db.Username, conf.Db.Password, conf.Db.Address, conf.Db.Name)

	r := gin.Default()
	// routes------------------------------------------
	r.GET("", root)
	// user============================================
	r.GET("users", controllers.GetUsers)
	r.POST("user", controllers.PostUser)
	// auth============================================
	private := r.Group("/private")
	private.Use(controllers.AuthRequired)
	{
		private.GET("/status", controllers.Status)
	}
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	//-------------------------------------------------

	if err := r.Run(":3000"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ROOT BLANK PAGE",
	})
}
