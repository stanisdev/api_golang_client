package middlewares

import (
	"app/services"
	"app/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	_ "fmt"
)

type Notification struct {
	Message string `json:"message"`
	Image string `json:"image"`
	Header string `json:"header"`
	Priority string `json:"priority"`
	Expired string `json:"expired"`
	Button string `json:"button"`
	Link string `json:"link"`
	Company string `json:"company"`
}

func ValidateNotification(c *gin.Context) {
	var ntf Notification
	c.BindJSON(&ntf)

	cmp := &models.Company{
		Name: ntf.Company,
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

	exp, err0 := time.Parse("2006/01/02", ntf.Expired) // Parse exppired data
	if (err0 != nil) {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	prior, err1 := strconv.ParseUint(ntf.Priority, 10, 32)
	if err1 != nil {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	ntfInstance := &models.Notification{ // Creating Notification
		Message: ntf.Message,
		Image: ntf.Image,
		Header: ntf.Header,
		Priority: uint(prior),
		Expired: exp,
		Button: ntf.Button,
		Link: ntf.Link,
		CompanyID: cmpId,
	}
	if (!models.ValidateModel(ntfInstance)) {
		services.WrongPostData(c)
		c.Abort()
		return
	}
	c.Set("notificationBlank", ntfInstance)
	c.Next()
}