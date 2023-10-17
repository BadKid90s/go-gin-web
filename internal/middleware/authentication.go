package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/errors"
	"go-gin-demo/internal/common/resp"
	"slices"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions := []string{
			"/system/user",
		}
		requestURI := c.Request.RequestURI

		if slices.Index(permissions, requestURI) >= 0 {
			c.Next()
		} else {
			resp.HandleError(c, errors.NewInternalError("没有权限！"), nil)
			c.Abort()
			return
		}
	}
}
