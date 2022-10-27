package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/database"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/models"
)

func GetAllPostsHandler(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	c.JSON(200, posts)
}

func CreatePostHandler(c *gin.Context) {
	var body struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  body.UserID,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "post created"})
}
