package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/database"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/models"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/usecases"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/utils"
)

var users []models.User

func GetAllUsersHandler(c *gin.Context) {
	var users []models.User
	if err := database.DB.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"data": utils.ToSerializedArray(users)})
}

func CreateUserHandler(c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	}

	id, err := usecases.CreateUserUseCase(user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user_id": id})
}
