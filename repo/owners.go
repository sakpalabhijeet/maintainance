package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type OwnerRepository interface {
	Create(ctx context.Context, owner *models.Owners) error
	GetByID(ctx context.Context, id int64) (*models.Owners, error)
	GetAll(ctx context.Context) ([]models.Owners, error)
	Update(ctx context.Context, owner *models.Owners) error
	Delete(ctx context.Context, id int64) error
}

type ownerRepository struct {
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnerRepository {
	return &ownerRepository{db: db}
}

func (r *ownerRepository) Create(ctx context.Context, owner *models.Owners) error {
	return r.db.WithContext(ctx).Create(owner).Error
}

func (r *ownerRepository) GetByID(ctx context.Context, id int64) (*models.Owners, error) {
	var owner models.Owners
	err := r.db.WithContext(ctx).First(&owner, id).Error
	if err != nil {
		return nil, err
	}
	return &owner, nil
}

func (r *ownerRepository) GetAll(ctx context.Context) ([]models.Owners, error) {
	var owners []models.Owners
	err := r.db.WithContext(ctx).Find(&owners).Error
	return owners, err
}

func (r *ownerRepository) Update(ctx context.Context, owner *models.Owners) error {
	return r.db.WithContext(ctx).Save(owner).Error
}

func (r *ownerRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Owners{}, id).Error
}
