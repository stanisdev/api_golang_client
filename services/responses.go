package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrongPostData(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"ok": false,
		"message": "Wrong post data",
	})
}

func WrongUrlParams(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"ok": false,
		"message": "Wrong URL params",
	})
}