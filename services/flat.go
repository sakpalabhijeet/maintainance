package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
	"errors"
)

type FlatService interface {
	CreateFlat(ctx context.Context, flat *models.Flat) error
}

type flatService struct {
	repo repo.FlatRepository
}

func NewFlatService(repo repo.FlatRepository) FlatService {
	return &flatService{repo: repo}
}

func (s *flatService) CreateFlat(ctx context.Context, flat *models.Flat) error {
	if flat.AreaSqFt <= 0 {
		return errors.New("Area must be greater than 0")
	}
	return s.repo.Create(ctx, flat)
}
