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
	r := engine()

	configs.Connect(conf.Db.Username, conf.Db.Password, conf.Db.Address, conf.Db.Name)
	r.Use(gin.Logger())

	if err := engine().Run(":3000"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	// routes------------------------------------------
	r.GET("", root)
	// user============================================
	r.GET("users", controllers.GetUsers)
	r.POST("user", controllers.PostUser)
	// auth============================================
	private := r.Group("/private")
	private.Use(controllers.AuthRequired)
	{
		private.GET("/me", controllers.Me)
		private.GET("/status", controllers.Status)
	}
	//-------------------------------------------------

	return r
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ROOT BLANK PAGE",
	})
}
