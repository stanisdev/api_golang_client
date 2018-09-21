package controllers

import "github.com/gin-gonic/gin"

func NotificationList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Notification List",
	})
}