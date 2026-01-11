package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type FlatRepository interface {
	Create(ctx context.Context, flat *models.Flat)error
}

type flatRepository struct{
	db *gorm.DB
}

func NewFlatReporsitry(db *gorm.DB) FlatRepository{
	return &flatRepository{db:db}
}

func (r *flatRepository)Create (ctx context.Context, flat *models.Flat)error{
	return r.db.WithContext(ctx).Create(flat).Error
}