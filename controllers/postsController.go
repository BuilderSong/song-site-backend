package controllers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Body     string
		Title    string
		Topic    string
		Image    string
		Abstract string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageB64data := body.Image[strings.IndexByte(body.Image, ',')+1:]

	fmt.Println(imageB64data)

	imageContent, err := base64.StdEncoding.DecodeString(imageB64data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create post and save to db

	post := models.Post{Title: body.Title, Body: body.Body, Topic: body.Topic, Image: imageContent, Abstract: body.Abstract}

	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
		"posts":   post})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")

	//get the posts
	var post models.Post
	result := initializers.DB.Find(&post, id)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body     string
		Title    string
		Topic    string
		Abstract string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.Find(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Topic: body.Topic, Abstract: body.Abstract})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "deleted ok",
	})
}
