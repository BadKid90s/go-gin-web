package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/system/request"
	"go-gin-demo/internal/system/service"
	"net/http"
)

type UserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UserInfo(ctx *gin.Context)
}

func NewUserHandler(handler *common.Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

type userHandler struct {
	*common.Handler
	userService service.UserService
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(request.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		common.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if err := h.userService.Register(ctx, req); err != nil {
		common.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	common.HandleSuccess(ctx, nil)
}

func (h *userHandler) UserInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"user": "golang body",
		"age":  18,
	})
}

func (h *userHandler) Login(ctx *gin.Context) {
	// Binding from JSON
	type Login struct {
		User     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var login Login
	if err := ctx.ShouldBind(&login); err == nil {
		fmt.Printf("login info:%#v\n", login)
		ctx.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
