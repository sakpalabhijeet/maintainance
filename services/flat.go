package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type FlatService interface {
	Create(ctx context.Context, flat *models.Flat) error
	GetByID(ctx context.Context, id int64) (*models.Flat, error)
	GetAll(ctx context.Context) ([]models.Flat, error)
	Update(ctx context.Context, flat *models.Flat) error
	Delete(ctx context.Context, id int64) error
}

type flatService struct {
	repo repo.FlatRepository
}

func NewFlatService(repo repo.FlatRepository) FlatService {
	return &flatService{repo: repo}
}

func (s *flatService) Create(ctx context.Context, flat *models.Flat) error {
	return s.repo.Create(ctx, flat)
}

func (s *flatService) GetByID(ctx context.Context, id int64) (*models.Flat, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *flatService) GetAll(ctx context.Context) ([]models.Flat, error) {
	return s.repo.GetAll(ctx)
}

func (s *flatService) Update(ctx context.Context, flat *models.Flat) error {
	return s.repo.Update(ctx, flat)
}

func (s *flatService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
