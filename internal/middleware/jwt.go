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
			common.HandleError(c, common.NewBarRequestError("jwt token 'Authorization' not found in header "), nil)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			common.HandleError(c, err, nil)
			c.Abort()
			return
		}
		c.Set(common.UserId, claims.UserId)
		c.Next()
	}
}
