package main

import (
	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DatabaseConnector()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Subscriber{})
}
