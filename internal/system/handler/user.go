package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/system/request"
	"go-gin-demo/internal/system/service"
	"go-gin-demo/pkg/jwt"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
}

func NewUserHandler(jwt *jwt.JWT, handler *common.Handler, userService service.UserService) UserHandler {
	return &userHandler{
		jwt:         jwt,
		Handler:     handler,
		userService: userService,
	}
}

type userHandler struct {
	*common.Handler
	jwt         *jwt.JWT
	userService service.UserService
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(request.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		common.HandleError(ctx, err, nil)
		return
	}
	if err := h.userService.Register(ctx, req); err != nil {
		common.HandleError(ctx, err, nil)
		return
	}
	common.HandleSuccess(ctx, nil)
}

func (h *userHandler) UserInfo(ctx *gin.Context) {
	userId := ctx.GetUint(common.UserId)
	userInfo, err := h.userService.UserInfo(ctx, userId)
	if err != nil {
		common.HandleError(ctx, err, nil)
		return
	}
	common.HandleSuccess(ctx, userInfo)
}

func (h *userHandler) Login(ctx *gin.Context) {
	loginReq := new(request.LoginRequest)
	if err := ctx.ShouldBind(&loginReq); err == nil {
		fmt.Printf("login info:%#v", loginReq)
	}
	user, err := h.userService.Login(ctx, loginReq)
	if err != nil {
		common.HandleError(ctx, err, nil)
		return
	}

	token, err := h.jwt.GenToken(user.Id)
	if err != nil {
		common.HandleError(ctx, err, nil)
		return
	}
	common.HandleSuccess(ctx, token)

}
