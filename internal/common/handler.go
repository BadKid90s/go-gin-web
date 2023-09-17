package common

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	//TODO jwt中提取userId
	return ""
}
