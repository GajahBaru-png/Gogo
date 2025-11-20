package main

import (
	"github.com/gin-gonic/gin"
	"gogo/controllers"
	"gogo/middleware"
	"gogo/models"
	"log"
	"github.com/joho/godotenv"
)	

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	router := gin.Default()
	
	models.ConnectDatabase()
	
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	router.GET("/login/:id", controllers.FindUser)
	router.GET("/items", controllers.FindItems)
	protected := router.Group("/admin").Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
		protected.POST("/items", controllers.Items)
		protected.PATCH("/items/:id", controllers.UpdateItems)
		protected.DELETE("/items/:id", controllers.DeleteItems)
	}
	router.Run(":8000")
}

