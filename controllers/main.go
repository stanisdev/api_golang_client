package controllers

import (
	"app/models"
	"app/middlewares"
	"app/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Env struct {
	db *gorm.DB
	DBMethods *models.DbMethods
}

func Start() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/uploads", services.GetDynamicConfig()["UploadsDir"])
	router.Use(middlewares.RequireAuthToken)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"ok": false,
			"message": "Page not found",
		})
	})
	env := &Env{
		db: models.GetConnection(),
		DBMethods: models.GetDmInstance(),
	}
	user := router.Group("/user")
	{
		user.POST("/login", env.UserLogin)
	}
	notification := router.Group("/notification")
	{
		notification.GET("/", env.NotificationList)
		notification.POST("/create", env.NotificationCreate)
		notification.DELETE("/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, env.NotificationRemove)
	}
	image := router.Group("/image")
	{
		image.POST("/upload", env.ImageUpload)
	}	
	router.Run(":" + viper.GetString("environment.port"))
}