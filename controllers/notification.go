package controllers

import "github.com/gin-gonic/gin"

func (e *Env) NotificationList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Notification List",
	})
}