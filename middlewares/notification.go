package middlewares

import (
	"app/models"
	"app/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindNotificationById(c *gin.Context) {
	id := c.MustGet("id").(uint)
	ntf := models.GetDmInstance().FindNotificationById(id)

	if (ntf.ID < 1) { // Notification not found
		services.WrongUrlParams(c)
		c.Abort()
		return
	}
	c.Set("notification", ntf)
	c.Set("oldCompanyID", ntf.CompanyID)
	c.Next()
}

func LimitOffset(c *gin.Context) {
	var lmt, ofst int
	limit := c.Query("limit")
	offset := c.Query("offset")
	if (len(offset) < 1) { // Offset
		ofst = 0
	} else {
		o, err := strconv.Atoi(offset)
		if err != nil || o < 0 {
			services.WrongUrlParams(c)
			c.Abort()
			return
		} else {
			ofst = o
		}
	}
	if (len(limit) < 1) { // Limit
		lmt = 10
	} else {
		i, err := strconv.Atoi(limit)
		if err != nil || i < 1 {
			services.WrongUrlParams(c)
			c.Abort()
			return
		} else {
			lmt = i
		}
	}
	c.Set("limit", lmt)
	c.Set("offset", ofst)
	c.Next()
}