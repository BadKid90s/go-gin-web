package routers

import (
	"github.com/gin-gonic/gin"
	middleware2 "go-gin-demo/internal/common/middleware"
	"go-gin-demo/internal/knowledge"
	"go-gin-demo/internal/system"
	"go-gin-demo/pkg/jwt"
	"net/http"
)

func NewRouter(
	jwt *jwt.JWT,
	r *gin.Engine,
	system *system.System,
	kd *knowledge.Knowledge,
) http.Handler {
	//中间件
	corsMiddleware := middleware2.Cors()
	jwtMiddleware := middleware2.Jwt(jwt)
	authMiddleware := middleware2.Authentication()

	api := r.RouterGroup
	api.Use(corsMiddleware)

	noAuthRouter := api.Group("/")
	noAuthRouter.POST("/register", system.User.Register)
	noAuthRouter.POST("/login", system.User.Login)

	systemApi := api.Group("/system").Use(jwtMiddleware, authMiddleware)
	systemApi.GET("/user", system.User.UserInfo)

	knowledgebaseApi := api.Group("/knowledgebase")
	knowledgebaseApi.POST("/collection", kd.CollectionHandler.Create)

	return r
}
