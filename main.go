package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/controllers"
	"github.com/rudolfoborges/go-gin-gorm-blog-api/database"
)

func init() {
	database.DatabaseInit()
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", controllers.GetAllUsersHandler)
		v1.POST("/users", controllers.CreateUserHandler)
		v1.POST("/posts", controllers.CreatePostHandler)
		v1.GET("/posts", controllers.GetAllPostsHandler)
	}

	router.Run(":3000")
}
