package repo

import (
	"Maintainance/models"
	"context"
	"gorm.io/gorm"
)

type OwnersRepository interface {
	Create(ctx context.Context, Owners *models.Owners)error
}

type ownersRepository struct{
	db *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) OwnersRepository{
	return &ownersRepository{db: db}
}

func (r *ownersRepository)Create (ctx context.Context, Owners *models.Owners)error{
	return r.db.WithContext(ctx).Create(Owners).Error
}