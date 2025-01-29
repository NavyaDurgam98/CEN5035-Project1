package controllers

import (
	"context"
	"net/http"
	"skillarcade-backend/models"
	"skillarcade-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid input"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into MongoDB
	collection := utils.GetCollection("users")
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "User registered successfully", "id": result.InsertedID})
}
