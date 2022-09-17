package main

import (
	"login-jwt/controllers"
	"login-jwt/db"
	"login-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.SetupDB()
	db.AutoMigrate(models.User{})
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}
	return router
}
