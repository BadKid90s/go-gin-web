package system

import (
	"go-gin-demo/internal/system/handler"
	"go-gin-demo/internal/system/repository"
	"go-gin-demo/internal/system/service"
)

var (
	NewUserHandler    = handler.NewUserHandler
	NewUserService    = service.NewUserService
	NewUserRepository = repository.NewUserRepository
)

type System struct {
	User handler.UserHandler
}

func NewSystem(uerHandler handler.UserHandler) *System {
	return &System{
		User: uerHandler,
	}
}
