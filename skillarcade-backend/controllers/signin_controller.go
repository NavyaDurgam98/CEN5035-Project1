package controllers

import (
	"context"
	"net/http"
	"skillarcade-backend/models"
	"skillarcade-backend/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	// Find user in MongoDB
	collection := utils.GetCollection("users")
	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"username": input.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "User not found"})
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid credentials"})
		return
	}

	// Generate auth token (dummy implementation)
	token := "dummy-auth-token"

	// Set cookie
	c.SetCookie("auth_token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"msg": "Sign-in successful", "token": token})
}
