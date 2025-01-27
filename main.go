package main

import (
	"os"

	routes "SkillArcade/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//post - signup
	//post signin

	router.POST("/api-1", func(c *gin.Context) {
		c.json(200, gin.H{"sucess": "Access Granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.json(200, gin.H{"sucess": "Access Granted for api-2"})

	})

	router.Run(":" + port)
}
