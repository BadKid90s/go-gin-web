package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/system/handler"
	"net/http"
)

func NewRouter(
	r *gin.Engine,
	systemHandler *handler.SystemHandler,
) http.Handler {

	api := r.RouterGroup

	noAuthRouter := api.Group("/")
	noAuthRouter.POST("/register", systemHandler.User.Register)
	noAuthRouter.POST("/login", systemHandler.User.Login)

	systemApi := api.Group("/system")
	systemApi.GET("/user", systemHandler.User.UserInfo)

	return r
}
