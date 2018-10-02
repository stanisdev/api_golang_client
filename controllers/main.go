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

// @TODO: Generate ID for notifications
func Start() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
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
		user.GET("/profile", env.UserProfile)
		user.POST("/change/password", env.UserChangePassword)
	}
	notification := router.Group("/notification")
	{
		notification.GET("/list", env.NotificationList)
		notification.POST("/create", middlewares.ValidateNotification, env.NotificationCreate)
		notification.GET("/delete/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, env.NotificationRemove)
		notification.GET("/get/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, env.NotificationGetById)
		notification.POST("/edit/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, middlewares.ValidateNotification, env.NotificationUpdate)
		notification.GET("/count", env.NotificationCount)
	}
	image := router.Group("/image")
	{
		image.POST("/upload", env.ImageUpload)
	}	
	router.Run(":" + viper.GetString("environment.port"))
}