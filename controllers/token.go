package controllers

import (
	"login-jwt/db"
	"login-jwt/helper"
	"login-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokeRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if email is existed and password is correct
	record := db.SetupDB().Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": record.Error.Error()})
		return
	}
	tokenString, err := helper.GenerateJWT(user.UserName, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
