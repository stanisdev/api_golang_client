package middlewares

import (
	"app/services"
	"app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func invalidToken(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"ok": false,
		"message": "Wrong Auth Token",
	})
	c.Abort()
}

func RequireAuthToken(c *gin.Context) {
	publicUrls := []string{ "/user/login" }
	currentUrl := c.Request.URL.Path
	for _, url := range publicUrls {
    if (url == currentUrl) {
			c.Next()
			return
		}
	}
	authHeader := c.Request.Header.Get("authorization")
	s := strings.Split(authHeader, "Bearer ")
	if (len(s) < 2) { // Token was not given
		invalidToken(c)
		return
	}
	token := strings.Trim(s[1], " ")
	userJWT, isValid := services.DecryptJWT(token)
	if (!isValid) { // Invalid token
		invalidToken(c)
		return
	}
	user := models.User{} // Attach user to context
	models.GetConnection().Where("ID = ?", userJWT.ID).First(&user)
	c.Set("user", user)
	c.Next()
}