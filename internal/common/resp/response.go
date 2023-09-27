package resp

import (
	sysErrors "errors"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/errors"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: errors.SuccessCode, Message: "", Data: data}
	ctx.JSON(http.StatusOK, resp)
}

func HandleError(ctx *gin.Context, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	var code int
	var sysErr *errors.SystemError
	switch {
	case sysErrors.As(err, &sysErr):
		code = sysErr.Code
	default:
		code = errors.UnknownErrorCode
	}
	resp := Response{Code: code, Message: err.Error(), Data: data}
	ctx.JSON(http.StatusOK, resp)
}
