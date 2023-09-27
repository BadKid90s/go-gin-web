package service

import (
	"context"
	"go-gin-demo/internal/common/errors"
	"go-gin-demo/internal/system/model"
	"go-gin-demo/internal/system/pkg"
	"go-gin-demo/internal/system/repository"
	"go-gin-demo/internal/system/request"
	"go-gin-demo/internal/system/response"
)

type UserService interface {
	Register(ctx context.Context, req *request.RegisterRequest) error
	Login(ctx context.Context, req *request.LoginRequest) (*model.User, error)
	UserInfo(ctx context.Context, userId uint) (*response.UserInfoResponse, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) UserInfo(ctx context.Context, userId uint) (*response.UserInfoResponse, error) {
	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}
	return &response.UserInfoResponse{
		UserName:  user.UserName,
		LoginName: user.LoginName,
		Email:     user.Email,
		Mobile:    user.Mobile,
	}, nil
}

func (s *userService) Register(ctx context.Context, req *request.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByLoginName(ctx, *req.LoginName); err == nil && user != nil {
		return errors.NewBizError("LoginName already exists")
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

func (s *userService) Login(ctx context.Context, req *request.LoginRequest) (*model.User, error) {
	user, err := s.userRepo.GetByLoginName(ctx, req.LoginName)
	if err != nil || user == nil {
		return nil, errors.NewBizError("failed to get user by loginName")
	}
	if !pkg.CheckPassword(req.Password, *user.Password) {
		return nil, errors.NewBizError("password verify failed")
	}
	return user, nil
}
