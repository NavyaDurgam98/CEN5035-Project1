package main

import (
	"skillarcade-backend/controllers"
	"skillarcade-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	utils.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/register", controllers.Register)
	router.POST("/signIn", controllers.SignIn)

	// Start the server
	router.Run(":8080")
}
