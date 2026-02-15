package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.Users) error
	GetByID(ctx context.Context, id uint) (*models.Users, error)
	GetByEmail(ctx context.Context, email string) (*models.Users, error)
	GetAll(ctx context.Context) ([]models.Users, error)
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.Users) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.Users, error) {
	var user models.Users
	err := r.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetAll(ctx context.Context) ([]models.Users, error) {
	var users []models.Users
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Users{}, id).Error
}
