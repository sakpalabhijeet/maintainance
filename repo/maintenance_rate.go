package repo

import (
	"Maintainance/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type MaintenanceRateRepository interface {
	Create(ctx context.Context, rate *models.MaintenanceRate) error
	GetByID(ctx context.Context, id int64) (*models.MaintenanceRate, error)
	GetAll(ctx context.Context) ([]models.MaintenanceRate, error)
	GetActiveRates(ctx context.Context, billMonth time.Time) ([]models.MaintenanceRate, error)
	Update(ctx context.Context, rate *models.MaintenanceRate) error
	Delete(ctx context.Context, id int64) error
}

type maintenanceRateRepository struct {
	db *gorm.DB
}

func NewMaintenanceRateRepository(db *gorm.DB) MaintenanceRateRepository {
	return &maintenanceRateRepository{db: db}
}

func (r *maintenanceRateRepository) Create(
	ctx context.Context,
	rate *models.MaintenanceRate,
) error {
	return r.db.WithContext(ctx).Create(rate).Error
}

func (r *maintenanceRateRepository) GetByID(
	ctx context.Context,
	id int64,
) (*models.MaintenanceRate, error) {

	var rate models.MaintenanceRate

	err := r.db.WithContext(ctx).
		First(&rate, id).Error

	if err != nil {
		return nil, err
	}

	return &rate, nil
}

func (r *maintenanceRateRepository) GetAll(
	ctx context.Context,
) ([]models.MaintenanceRate, error) {

	var rates []models.MaintenanceRate

	err := r.db.WithContext(ctx).
		Order("start_date desc").
		Find(&rates).Error

	return rates, err
}

func (r *maintenanceRateRepository) GetActiveRates(
	ctx context.Context,
	billMonth time.Time,
) ([]models.MaintenanceRate, error) {

	var rates []models.MaintenanceRate

	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Where("start_date <= ?", billMonth).
		Where("(end_date IS NULL OR end_date >= ?)", billMonth).
		Find(&rates).Error

	return rates, err
}

func (r *maintenanceRateRepository) Update(
	ctx context.Context,
	rate *models.MaintenanceRate,
) error {

	return r.db.WithContext(ctx).
		Save(rate).Error
}

func (r *maintenanceRateRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	return r.db.WithContext(ctx).
		Delete(&models.MaintenanceRate{}, id).Error
}
