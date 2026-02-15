package repo

import (
	"Maintainance/models"
	"context"
	"gorm.io/gorm"
)

type FlatRepository interface {
	Create(ctx context.Context, flat *models.Flat) error
	GetByID(ctx context.Context, id int64) (*models.Flat, error)
	GetAll(ctx context.Context) ([]models.Flat, error)
	Update(ctx context.Context, flat *models.Flat) error
	Delete(ctx context.Context, id int64) error
}

type flatRepository struct {
	db *gorm.DB
}

func NewFlatRepository(db *gorm.DB) FlatRepository {
	return &flatRepository{db: db}
}

func (r *flatRepository) Create(ctx context.Context,flat *models.Flat,) error {
	return r.db.WithContext(ctx).Create(flat).Error
}

func (r *flatRepository) GetByID(ctx context.Context,id int64,) (*models.Flat, error) {

	var flat models.Flat

	err := r.db.WithContext(ctx).
		First(&flat, id).Error

	if err != nil {
		return nil, err
	}

	return &flat, nil
}


func (r *flatRepository) GetAll(ctx context.Context,) ([]models.Flat, error) {

	var flats []models.Flat

	err := r.db.WithContext(ctx).
		Find(&flats).Error

	return flats, err
}

func (r *flatRepository) Update(ctx context.Context,flat *models.Flat,) error {

	return r.db.WithContext(ctx).
		Model(&models.Flat{}).
		Where("id = ?", flat.Id).
		Updates(map[string]interface{}{
			"flat_number": flat.FlatNumber,
			"owner_id":    flat.OwnerId,
			"area_sq_ft":  flat.AreaSqFt,
			"is_occupied": flat.IsOccupied,
		}).Error
}
func (r *flatRepository) Delete(ctx context.Context,id int64,) error {

	return r.db.WithContext(ctx).
		Delete(&models.Flat{}, id).Error
}
