package controllers

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()
	user := router.Group("/user")
	{
		user.GET("/login", UserLogin)
	}
	notification := router.Group("/notification")
	{
		notification.GET("/", NotificationList)
	}
	router.Run()
}