package controllers

import (
	"app/models"
	"github.com/gin-gonic/gin"
	_ "fmt"
)

func (e *Env) NotificationList(c *gin.Context) {
	text := c.Query("text")
	ntfs := models.GetDmInstance().FindNotifications(text) // @TODO: Remove excessive fields
	c.JSON(200, gin.H{
		"ok": true,
		"payload": ntfs,
	})
}

func (e *Env) NotificationCreate(c *gin.Context) {
	ntf := c.MustGet("notificationBlank").(*models.Notification)

	e.db.Create(ntf)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"id": ntf.ID,
			"message": ntf.Message,
			"image": ntf.Image,
			"header": ntf.Header,
			"priority": ntf.Priority,
			"expired": ntf.GetExpired(),
			"button": ntf.Button,
			"link": ntf.Link,
			"company": c.PostForm("company"),
			"created_at": ntf.CreatedAt.Unix(),
		},
	})
}

func (e *Env) NotificationUpdate(c *gin.Context) {
	ntfQuery := c.MustGet("notification").(*models.NotificationQuery)
	ntfBlank := c.MustGet("notificationBlank").(*models.Notification)
	ntfBlank.ID = ntfQuery.ID
	ntfBlank.CreatedAt = ntfQuery.CreatedAt 

	e.db.Save(&ntfBlank)

	oldCompanyID := c.MustGet("oldCompanyID").(uint)
	var count int
	e.db.Model(&models.Notification{}).Where("company_id = ?", oldCompanyID).Count(&count)
	if (count < 1) {
		e.db.Where("id = ?", oldCompanyID).Limit(1).Unscoped().Delete(&models.Company{}) // Remove outdated single company
	}
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"id": ntfBlank.ID,
			"message": ntfBlank.Message,
			"image": ntfBlank.Image,
			"header": ntfBlank.Header,
			"priority": ntfBlank.Priority,
			"expired": ntfBlank.GetExpired(),
			"button": ntfBlank.Button,
			"link": ntfBlank.Link,
			"company": c.PostForm("company"),
			"created_at": ntfBlank.CreatedAt.Unix(),
		},
	})
}

func (e *Env) NotificationRemove(c *gin.Context) {
	ntf := c.MustGet("notification").(*models.NotificationQuery)
	e.db.Where("id = ?", ntf.ID).Limit(1).Unscoped().Delete(&models.Notification{})

	var count int
	e.db.Model(&models.Notification{}).Where("company_id = ?", ntf.CompanyID).Count(&count)
	if (count < 1) {
		e.db.Where("id = ?", ntf.CompanyID).Limit(1).Unscoped().Delete(&models.Company{})
	}
	c.JSON(200, gin.H{
		"ok": true,
	})
}

func (e *Env) NotificationGetById(c *gin.Context) {
	ntf := c.MustGet("notification").(*models.NotificationQuery)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": gin.H{
			"id": ntf.ID,
			"message": ntf.Message,
			"image": ntf.Image,
			"header": ntf.Header,
			"company": ntf.Company,
			"priority": ntf.Priority,
			"expired": ntf.GetExpired(),
			"button": ntf.Button,
			"link": ntf.Link,
			"created_at": ntf.CreatedAt.Unix(),
		},
	})
}