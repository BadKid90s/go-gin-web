package repository

import (
	"context"
	"github.com/pkg/errors"
	"go-gin-demo/internal/common"
	"go-gin-demo/internal/system/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

func NewUserRepository(r *common.Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*common.Repository
}

func (r *userRepository) Create(_ context.Context, user *model.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) GetByUsername(_ context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("user_name = ? ", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}
	return &user, nil
}
