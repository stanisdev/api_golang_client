package controllers

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"time"
)

func (e *Env) UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if (len(username) < 1 || len(password) < 1) {
		c.JSON(200, gin.H{
			"ok": false,
			"errors": gin.H{
				"username": "Enter password/username",
			},
		})
		return
	}
	user, exists := e.DBMethods.UserFindOne(&models.User{Username: username})
	wrongCred := func () {
		c.JSON(200, gin.H{
			"ok": false,
			"errors": gin.H{
				"username": "Wrong password/username",
			},
		})
	}
	if (!exists) {
		wrongCred()
		return
	}
	isValid := services.CheckPasswordHash(password + user.Salt, user.Password)
	if (!isValid) {
		wrongCred()
		return
	}
	tokenString, err := services.CryptJWT(user.ID)
	if (err != nil) {
		c.JSON(200, gin.H{
			"ok": false,
			"message": "Can not log in",
		})
		return
	}
	user.LastLogin = time.Now()
	e.db.Save(&user)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"id_token": tokenString,
		},
	})
}