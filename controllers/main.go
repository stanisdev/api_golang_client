package controllers

import (
	"app/models"
	"app/middlewares"
	"app/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "fmt"
)

type Env struct {
	db *gorm.DB
	DBMethods *models.DbMethods
}

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
	prefix := viper.GetString("environment.prefix")
	user := router.Group(prefix + "/user")
	{
		user.POST("/login", env.UserLogin)
		user.GET("/profile", env.UserProfile)
		user.POST("/change/password", env.UserChangePassword)
	}
	notification := router.Group(prefix + "/notification")
	{
		notification.GET("/list", middlewares.LimitOffset, env.NotificationList)
		notification.POST("/create", middlewares.ValidateNotification, env.NotificationCreate)
		notification.GET("/delete/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, env.NotificationRemove)
		notification.GET("/get/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, env.NotificationGetById)
		notification.POST("/edit/:id", middlewares.UrlIdCorrectness, middlewares.FindNotificationById, middlewares.ValidateNotification, env.NotificationUpdate)
		notification.GET("/count", env.NotificationCount)
	}
	image := router.Group(prefix + "/image")
	{
		image.POST("/upload", env.ImageUpload)
	}	
	router.GET("/notifications", env.NotificationPublic)
	router.Run(":" + viper.GetString("environment.port"))
}