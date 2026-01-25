package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type OwnerService interface {
	CreateOwner(ctx context.Context, owners *models.Owners)error
}

type ownerService struct{
	repo repo.OwnersRepository
}

func NewOwnerService ( repo repo.OwnersRepository) OwnerService {
	return &ownerService{repo: repo}
}

func (s *ownerService) CreateOwner (ctx context.Context, owner *models.Owners)error{
	return s.repo.Create(ctx, owner)
}