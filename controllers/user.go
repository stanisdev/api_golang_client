package controllers

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"time"
	_ "fmt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e *Env) UserLogin(c *gin.Context) {
	var login Login
	c.BindJSON(&login)
		
	if (len(login.Username) < 1 || len(login.Password) < 1) {
		services.WrongPostData(c)
		return
	}
	user, exists := e.DBMethods.UserFindOne(&models.User{Username: login.Username})
	wrongCred := func () {
		c.JSON(200, gin.H{
			"ok": false,
			"errors": gin.H{
				"pass": "Wrong password/username",
			},
		})
	}
	if (!exists) {
		wrongCred()
		return
	}
	isValid := services.CheckPasswordHash(login.Password + user.Salt, user.Password)
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

func (e *Env) UserProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"username": user.Username,
			"created_ad": user.CreatedAt.Unix(),
			"last_login": user.LastLogin.Unix(),
			"password_changed": user.PasswordChanged.Unix(),
		},
	})
}

func (e *Env) UserChangePassword(c *gin.Context) {
	oldPass := c.PostForm("old_password")
	newPass := c.PostForm("new_password")
	if (len(oldPass) < 1 || len(newPass) < 1 || oldPass == newPass) {
		services.WrongPostData(c)
		return
	}
	user := c.MustGet("user").(models.User)
	isValid := services.CheckPasswordHash(oldPass + user.Salt, user.Password)
	if (!isValid) {
		c.JSON(200, gin.H{
			"ok": false,
			"errors": gin.H{
				"old_password": "Wrong password",
			},
		})
		return
	}
	salt := services.GenerateRandomString(10)
	hash, err := services.GetPasswordHash(newPass + salt)
	if (err != nil) {
		services.ServerError(err, c)
		return
	}
	user.Salt = salt
	user.Password = hash
	user.PasswordChanged = time.Now()
	e.db.Save(&user)
	
	c.JSON(200, gin.H{
		"ok": true,
	})
}