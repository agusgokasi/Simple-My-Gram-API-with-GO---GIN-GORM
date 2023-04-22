package middleware

import (
	"project-final/helper"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helper.VerifyToken(c)
		if err != nil {
			helper.Unauthorized(c, err.Error())
			c.Abort()
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}
