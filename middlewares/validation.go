package middlewares

import (
	"app/services"
	"app/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func ValidateNotification(c *gin.Context) {
	cmp := &models.Company{
		Name: c.PostForm("company"),
	}
	if (!models.ValidateModel(cmp)) {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	exCmp := models.Company{}
	models.GetConnection().Where("name = ?", cmp.Name).First(&exCmp)

	var cmpId uint
	if (exCmp.ID < 1) { // No such company, let's create it
		models.GetConnection().Create(&cmp)
		cmpId = cmp.ID
	} else {
		cmpId = exCmp.ID
	}

	exp, err0 := time.Parse("2006/01/02", c.PostForm("expired")) // Parse exppired data
	if (err0 != nil) {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	prior, err1 := strconv.ParseUint(c.PostForm("priority"), 10, 32)
	if err1 != nil {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	ntf := &models.Notification{ // Creating Notification
		Message: c.PostForm("message"),
		Image: c.PostForm("image"),
		Header: c.PostForm("header"),
		Priority: uint(prior),
		Expired: exp,
		Button: c.PostForm("button"),
		Link: c.PostForm("link"),
		CompanyID: cmpId,
	}
	if (!models.ValidateModel(ntf)) {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	c.Set("notificationBlank", ntf)
	c.Next()
}