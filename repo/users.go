package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, users *models.Users)error
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB)UserRepository{
	return &userRepository{db: db}
}

func (r *userRepository)Create(ctx context.Context, User *models.Users)error{
	return r.db.WithContext(ctx).Create(User).Error
}