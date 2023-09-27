package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/constant"
	"go-gin-demo/internal/common/errors"
	"go-gin-demo/internal/common/resp"
	"go-gin-demo/pkg/jwt"
)

func Jwt(jwt *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			resp.HandleError(c, errors.NewBarRequestError("jwt token 'Authorization' not found in header "), nil)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			resp.HandleError(c, err, nil)
			c.Abort()
			return
		}
		c.Set(constant.UserId, claims.UserId)
		c.Next()
	}
}
