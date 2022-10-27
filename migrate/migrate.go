package main

import (
	"github.com/rudolfoborges/go-gin-gorm-blog-api/database"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/models"
)

func init() {
	database.DatabaseInit()
}

func main() {
	database.DB.AutoMigrate(&models.User{}, &models.Post{})
}
