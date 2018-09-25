package controllers

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"strconv"
	_ "fmt"
)

func (e *Env) NotificationList(c *gin.Context) {
	// user := c.MustGet("user").(models.User)
	c.JSON(200, gin.H{
		"ok": true,
		"message": "Notification List",
	})
}

func (e *Env) NotificationCreate(c *gin.Context) {
	cmp := &models.Company{
		Name: c.PostForm("company"),
	}
	if (!models.ValidateModel(cmp)) {
		services.WrongPostData(c)
		return
	}
	exCmp := models.Company{}
	e.db.Where("name = ?", cmp.Name).First(&exCmp)

	var cmpId uint
	if (exCmp.ID < 1) { // No such company, let's create it
		e.db.Create(&cmp)
		cmpId = cmp.ID
	} else {
		cmpId = exCmp.ID
	}

	ntf := &models.Notification{
		Text: c.PostForm("text"),
		Image: c.PostForm("image"),
		CompanyID: cmpId,
	}
	if (!models.ValidateModel(ntf)) {
		services.WrongPostData(c)
		return
	}
	e.db.Create(&ntf)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"id": ntf.ID,
			"text": ntf.Text,
			"image": ntf.Image,
			"company": cmp.Name,
		},
	})
}

func (e *Env) NotificationRemove(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if (err != nil) {
		services.WrongUrlParams(c)
		return
	}
	id := uint(u64)
	c.JSON(200, gin.H{
		"ok": true,
		"id": uint(id),
	})
}