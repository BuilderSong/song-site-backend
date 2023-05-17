package main

import (
	"os"
	"time"

	"github.com/BuilderSong/gin-json-crud/controllers"
	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DatabaseConnector()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("DOMAIN")},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/posts", middleware.AuthRequired, controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", middleware.AuthRequired, controllers.PostsUpdate)
	r.DELETE("/posts/:id", middleware.AuthRequired, controllers.PostsDelete)
	// r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.AuthRequired, controllers.Validate)
	r.POST("/logout", middleware.AuthRequired, controllers.Logout)
	r.POST("/sendEmails", middleware.AuthRequired, controllers.SendEmails)
	r.POST("/subscribe", controllers.Subscribe)
	r.PATCH("/subscribe", controllers.UnSubscribe)
	r.Run()
}
