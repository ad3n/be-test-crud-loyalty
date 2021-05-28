package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("api-key")
		if apiKey != os.Getenv("APP_KEY") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"message": "Apikey tidak valid",
			})

			return
		}

		c.Next()
	}
}
