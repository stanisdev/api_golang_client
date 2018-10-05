package controllers

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	_ "fmt"
)

type NotificationList struct {
	Id uint `json:"id"`
	Image string `json:"image"`
	Message string `json:"message"`
	Header string `json:"header"`
	Priority uint `json:"priority"`
	Expired string `json:"expired"`
	Button string `json:"button"`
	Link string `json:"link"`
	Company string `json:"company"`
}

func (e *Env) NotificationList(c *gin.Context) {
	text := c.Query("text")
	var lmt, ofst int
	limit := c.Query("limit") // @TODO: Move to middleware
	offset := c.Query("offset")
	if (len(offset) < 1) {
		ofst = 0
	} else {
		o, err := strconv.Atoi(offset)
		if err != nil || o < 0 {
			services.WrongUrlParams(c)
			return
		} else {
			ofst = o
		}
	}
	if (len(limit) < 1) {
		lmt = 10
	} else {
		i, err := strconv.Atoi(limit)
		if err != nil || i < 1 {
			services.WrongUrlParams(c)
			return
		} else {
			lmt = i
		}
	}
	ntfs := models.GetDmInstance().FindNotifications(text, lmt, ofst)
	var result []NotificationList
	var msg string
	for _, ntf := range *ntfs {
		msg = ntf.Message
		if (len(msg) > 62) {
			msg = msg[0:62] + "..."
		}
		result = append(result, NotificationList {
			Id: ntf.ID,
			Message: msg,
			Expired: ntf.Expired.Format("Jan 2 2006"),
			Link: ntf.Link,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"payload": result,
	})
}

func (e *Env) NotificationCreate(c *gin.Context) {
	ntf := c.MustGet("notificationBlank").(*models.Notification)
	e.db.Create(ntf)
	c.JSON(200, gin.H{
		"ok": true,
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
			"expired": ntf.Expired.Unix(),
			"button": ntf.Button,
			"link": ntf.Link,
			"created_at": ntf.CreatedAt.Unix(),
		},
	})
}

func (e *Env) NotificationCount(c *gin.Context) {
	var count int
	e.db.Model(&models.Notification{}).Count(&count)
	c.JSON(200, gin.H{
		"ok": true,
		"payload": count,
	})
}

func (e *Env) NotificationPublic(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok": true,
	})
}