package controllers

import "github.com/gin-gonic/gin"

func UserLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User Login",
	})
}