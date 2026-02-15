package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type UserService interface {
	Create(ctx context.Context, user *models.Users) error
	GetByID(ctx context.Context, id uint) (*models.Users, error)
	GetAll(ctx context.Context) ([]models.Users, error)
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, user *models.Users) error {
	return s.repo.Create(ctx, user)
}

func (s *userService) GetByID(ctx context.Context, id uint) (*models.Users, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) GetAll(ctx context.Context) ([]models.Users, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
