package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Env struct {
	db *gorm.DB
	DBMethods *models.DbMethods
}

func Start() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"ok": false,
			"message": "Page not found",
		})
	})
	env := &Env{
		db: models.GetConnection(), 
		DBMethods: &models.DbMethods{},
	}
	user := router.Group("/user")
	{
		user.POST("/login", env.UserLogin)
	}
	notification := router.Group("/notification")
	{
		notification.GET("/", env.NotificationList)
	}
	router.Run()
}