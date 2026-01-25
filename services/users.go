package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, users *models.Users)error
}

type userService struct{
	repo repo.UserRepository
}

func NewUserService ( repo repo.UserRepository) UserService{
	return &userService{repo: repo}
}

func (s *userService)CreateUser (ctx context.Context, users *models.Users)error{
	return s.repo.Create(ctx, users)
}