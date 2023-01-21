package controllers

import (
	"marcellinuselbert-be/initializers"
	"marcellinuselbert-be/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	if body.Title == "" && body.Body == "" {
		c.JSON(200, gin.H{
			"message": "No args",
		})
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsFetch(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsFetchById(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"posts": post,
	})
}
