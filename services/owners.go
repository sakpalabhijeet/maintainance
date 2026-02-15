package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type OwnerService interface {
	Create(ctx context.Context, owner *models.Owners) error
	GetByID(ctx context.Context, id int64) (*models.Owners, error)
	GetAll(ctx context.Context) ([]models.Owners, error)
	Update(ctx context.Context, owner *models.Owners) error
	Delete(ctx context.Context, id int64) error
}

type ownerService struct {
	repo repo.OwnerRepository
}

func NewOwnerService(r repo.OwnerRepository) OwnerService {
	return &ownerService{
		repo: r,
	}
}

func (s *ownerService) Create(ctx context.Context, owner *models.Owners) error {
	return s.repo.Create(ctx, owner)
}

func (s *ownerService) GetByID(ctx context.Context, id int64) (*models.Owners, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ownerService) GetAll(ctx context.Context) ([]models.Owners, error) {
	return s.repo.GetAll(ctx)
}

func (s *ownerService) Update(ctx context.Context, owner *models.Owners) error {
	return s.repo.Update(ctx, owner)
}

func (s *ownerService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
