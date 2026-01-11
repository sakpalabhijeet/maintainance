package repo

import (
	"Maintainance/models"
	"context"
	"gorm.io/gorm"
)

type SocietyRepository interface {
	Create(ctx context.Context, society *models.Society)error
}

type societiesRepository struct{
	db *gorm.DB
}

func NewSocietyRepository(db *gorm.DB) SocietyRepository{
	return &societiesRepository{db: db}
}

func (r *societiesRepository)Create (ctx context.Context, society *models.Society)error{
	return r.db.WithContext(ctx).Create(society).Error
}
