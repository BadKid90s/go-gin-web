package repository

import (
	"context"
	"github.com/pkg/errors"
	"go-gin-demo/internal/common/repository"
	"go-gin-demo/internal/system/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByLoginName(ctx context.Context, loginName string) (*model.User, error)
	GetById(ctx context.Context, id uint) (*model.User, error)
}

func NewUserRepository(r *repository.Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*repository.Repository
}

func (r *userRepository) GetById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	if err := r.DB.WithContext(ctx).Where("id = ? ", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by id")
	}
	return &user, nil
}

func (r *userRepository) GetByLoginName(ctx context.Context, loginName string) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Where("login_name = ? ", loginName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by loginName")
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.DB.WithContext(ctx).Where("user_name = ? ", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}
	return &user, nil
}
