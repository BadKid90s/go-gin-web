package service

import (
	"context"
	"github.com/pkg/errors"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/system/model"
	"go-gin-demo/internal/system/repository"
	"go-gin-demo/internal/system/request"
)

type UserService interface {
	Register(ctx context.Context, req *request.RegisterRequest) error
	//Login(ctx context.Context, req *request.LoginRequest) (string, error)
	//GetProfile(ctx context.Context, userId string) (*model.User, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Register(ctx context.Context, req *request.RegisterRequest) error {
	// check username
	if user, err := s.userRepo.GetByUsername(ctx, *req.Username); err == nil && user != nil {
		return common.ErrUsernameAlreadyUsed
	}
	// Create a user
	user := &model.User{
		UserName: *req.Username,
		Mobile:   *req.Mobile,
		Password: *req.Password,
		Status:   1,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

//
//func (s *userService) Login(ctx context.Context, req *request.LoginRequest) (string, error) {
//	user, err := s.userRepo.GetByUsername(ctx, req.Username)
//	if err != nil || user == nil {
//		return "", response.ErrUnauthorized
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
//	if err != nil {
//		return "", errors.Wrap(err, "failed to hash password")
//	}
//	token, err := s.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
//	if err != nil {
//		return "", errors.Wrap(err, "failed to generate JWT token")
//	}
//
//	return token, nil
//}
//
//func (s *userService) GetProfile(ctx context.Context, userId string) (*model.User, error) {
//	user, err := s.userRepo.GetByID(ctx, userId)
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to get user by ID")
//	}
//
//	return user, nil
//}
//
//func (s *userService) UpdateProfile(ctx context.Context, userId string, req *request.UpdateProfileRequest) error {
//	user, err := s.userRepo.GetByID(ctx, userId)
//	if err != nil {
//		return errors.Wrap(err, "failed to get user by ID")
//	}
//
//	user.Email = req.Email
//	user.Nickname = req.Nickname
//
//	if err = s.userRepo.Update(ctx, user); err != nil {
//		return errors.Wrap(err, "failed to update user")
//	}
//
//	return nil
//}
