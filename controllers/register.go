package controllers

import (
	"login-jwt/db"
	"login-jwt/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	record := db.SetupDB().Create(&user)
	if record.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": record.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user_id":   user.ID,
		"email":     user.Email,
		"user_name": user.Name,
	})
}
