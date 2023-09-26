package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/middleware"
	"go-gin-demo/internal/system/handler"
	"go-gin-demo/pkg/jwt"
	"net/http"
)

func NewRouter(
	jwt *jwt.JWT,
	r *gin.Engine,
	systemHandler handler.SystemHandler,
) http.Handler {

	api := r.RouterGroup
	api.Use(middleware.Cors())

	noAuthRouter := api.Group("/")
	noAuthRouter.POST("/register", systemHandler.User.Register)
	noAuthRouter.POST("/login", systemHandler.User.Login)

	jwtMiddleware := middleware.Jwt(jwt)

	systemApi := api.Group("/system").Use(jwtMiddleware)
	systemApi.GET("/user", systemHandler.User.UserInfo)

	return r
}
