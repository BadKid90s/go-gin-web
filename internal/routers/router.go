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

	api := r.RouterGroup
	api.Use(middleware.Cors())

	noAuthRouter := api.Group("/")
	noAuthRouter.POST("/register", system.User.Register)
	noAuthRouter.POST("/login", system.User.Login)

	jwtMiddleware := middleware.Jwt(jwt)

	systemApi := api.Group("/system").Use(jwtMiddleware)
	systemApi.GET("/user", system.User.UserInfo)

	return r
}
