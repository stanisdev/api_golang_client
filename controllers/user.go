package controllers

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"time"
	_ "github.com/jinzhu/gorm"
	_ "fmt"
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
	user.LastLogin = time.Now()
	e.db.Save(&user)
	c.JSON(200, gin.H{
		"message": user.Password,
		"username": isValid,
	})
}