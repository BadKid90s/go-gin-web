package service

import (
	"context"
	"github.com/pkg/errors"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/system/model"
	"go-gin-demo/internal/system/pkg"
	"go-gin-demo/internal/system/repository"
	"go-gin-demo/internal/system/request"
	"go-gin-demo/pkg/jwt"
)

type UserService interface {
	Register(ctx context.Context, req *request.RegisterRequest) error
	Login(ctx context.Context, req *request.LoginRequest) (string, error)
}

func NewUserService(jwt *jwt.JWT, userRepo repository.UserRepository) UserService {
	return &userService{
		jwt:      jwt,
		userRepo: userRepo,
	}
}

type userService struct {
	jwt      *jwt.JWT
	userRepo repository.UserRepository
}

func (s *userService) Register(ctx context.Context, req *request.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByLoginName(ctx, *req.LoginName); err == nil && user != nil {
		return common.NewBizError("LoginName already exists")
	}

	password, err := pkg.HashPassword(*req.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	// Create a user
	user := &model.User{
		UserName:  req.UserName,
		LoginName: req.LoginName,
		Password:  &password,
		Enabled:   true,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (s *userService) Login(ctx context.Context, req *request.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByLoginName(ctx, req.LoginName)
	if err != nil || user == nil {
		return "", common.NewBizError("failed to get user by loginName")
	}

	if !pkg.CheckPassword(req.Password, *user.Password) {
		return "", common.NewBizError("password verify failed")
	}

	token, err := s.jwt.GenToken(user.Id)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}
	return token, nil
}
