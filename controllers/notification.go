package controllers

import (
	// "app/models"
	"github.com/gin-gonic/gin"
)

func (e *Env) NotificationList(c *gin.Context) {
	// user := c.MustGet("user").(models.User)
	c.JSON(200, gin.H{
		"message": "Notification List",
	})
}