package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common"
	"go-gin-demo/pkg/jwt"
)

func Jwt(jwt *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			common.HandleError(c, common.NewBarRequestError("jwt not found"), nil)
			c.Abort()
			return
		}
		_, err := jwt.ParseToken(tokenString)
		if err != nil {
			common.HandleError(c, err, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
