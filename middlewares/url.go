package middlewares

import (
	"app/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UrlIdCorrectness(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if (err != nil) {
		services.WrongUrlParams(c)
		c.Abort()
		return
	}
	id := uint(u64)
	c.Set("id", id)
	c.Next()
}