package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
	"errors"
	"time"
)

type MaintenanceRateService interface {
	Create(ctx context.Context, rate *models.MaintenanceRate) error
	GetByID(ctx context.Context, id int64) (*models.MaintenanceRate, error)
	GetAll(ctx context.Context) ([]models.MaintenanceRate, error)
	GetActiveRates(ctx context.Context, billMonth time.Time) ([]models.MaintenanceRate, error)
	Update(ctx context.Context, rate *models.MaintenanceRate) error
	Delete(ctx context.Context, id int64) error
}

type maintenanceRateService struct {
	repo repo.MaintenanceRateRepository
}

func NewMaintenanceRateService(repo repo.MaintenanceRateRepository) MaintenanceRateService {
	return &maintenanceRateService{repo: repo}
}

func (s *maintenanceRateService) Create(
	ctx context.Context,
	rate *models.MaintenanceRate,
) error {

	if rate.RatePerSqFt <= 0 {
		return errors.New("rate_per_sq_ft must be greater than zero")
	}

	if rate.GSTPercent < 0 {
		return errors.New("gst_percent cannot be negative")
	}

	return s.repo.Create(ctx, rate)
}

func (s *maintenanceRateService) GetByID(
	ctx context.Context,
	id int64,
) (*models.MaintenanceRate, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *maintenanceRateService) GetAll(
	ctx context.Context,
) ([]models.MaintenanceRate, error) {
	return s.repo.GetAll(ctx)
}

func (s *maintenanceRateService) GetActiveRates(
	ctx context.Context,
	billMonth time.Time,
) ([]models.MaintenanceRate, error) {
	return s.repo.GetActiveRates(ctx, billMonth)
}

func (s *maintenanceRateService) Update(
	ctx context.Context,
	rate *models.MaintenanceRate,
) error {

	if rate.RatePerSqFt <= 0 {
		return errors.New("rate_per_sq_ft must be greater than zero")
	}

	return s.repo.Update(ctx, rate)
}

func (s *maintenanceRateService) Delete(
	ctx context.Context,
	id int64,
) error {
	return s.repo.Delete(ctx, id)
}
