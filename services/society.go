package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type SocietyService interface {
	CreateSociety(ctx context.Context, society *models.Society)error
}

type societyService struct{
	repo repo.SocietyRepository
}

func NewSocietyService(repo repo.SocietyRepository)SocietyService{
	return &societyService {repo: repo}
}

func (s *societyService)CreateSociety(ctx context.Context, society *models.Society)error{
	return s.repo.Create(ctx, society)
}