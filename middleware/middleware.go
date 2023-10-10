package middleware

import (
	"deliveryProduct/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SECRET = "secret"
)

func AuthValid(c *gin.Context) {
	err := token.TokenValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
