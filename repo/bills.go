package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(ctx context.Context, bill *models.Bill)error
}

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository{
	return &billRepository{db:db}
}

func (r *billRepository)Create (ctx context.Context, bill *models.Bill)error{
	return r.db.WithContext(ctx).Create(bill).Error
}