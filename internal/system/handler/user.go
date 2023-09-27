package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common/constant"
	"go-gin-demo/internal/common/resp"
	"go-gin-demo/internal/system/request"
	"go-gin-demo/internal/system/service"
	"go-gin-demo/pkg/jwt"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
}

func NewUserHandler(jwt *jwt.JWT, userService service.UserService) UserHandler {
	return &userHandler{
		jwt:         jwt,
		userService: userService,
	}
}

type userHandler struct {
	jwt         *jwt.JWT
	userService service.UserService
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(request.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	if err := h.userService.Register(ctx, req); err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) UserInfo(ctx *gin.Context) {
	userId := ctx.GetUint(constant.UserId)
	userInfo, err := h.userService.UserInfo(ctx, userId)
	if err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	resp.HandleSuccess(ctx, userInfo)
}

func (h *userHandler) Login(ctx *gin.Context) {
	loginReq := new(request.LoginRequest)
	if err := ctx.ShouldBind(&loginReq); err == nil {
		fmt.Printf("login info:%#v", loginReq)
	}
	user, err := h.userService.Login(ctx, loginReq)
	if err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}

	token, err := h.jwt.GenToken(user.Id)
	if err != nil {
		resp.HandleError(ctx, err, nil)
		return
	}
	resp.HandleSuccess(ctx, token)

}
