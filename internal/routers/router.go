package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/middleware"
	"go-gin-demo/internal/system"
	"go-gin-demo/pkg/jwt"
	"net/http"
)

func NewRouter(
	jwt *jwt.JWT,
	r *gin.Engine,
	system *system.System,
) http.Handler {
	//中间件
	corsMiddleware := middleware.Cors()
	jwtMiddleware := middleware.Jwt(jwt)
	authMiddleware := middleware.Authentication()

	api := r.RouterGroup
	api.Use(corsMiddleware)

	noAuthRouter := api.Group("/")
	noAuthRouter.POST("/register", system.User.Register)
	noAuthRouter.POST("/login", system.User.Login)

	systemApi := api.Group("/system").Use(jwtMiddleware, authMiddleware)
	systemApi.GET("/user", system.User.UserInfo)

	return r
}
