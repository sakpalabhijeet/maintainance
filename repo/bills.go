package repo

import (
	"Maintainance/models"
	"context"

	"gorm.io/gorm"
)

type BillRepository interface {
	Create(ctx context.Context, bill *models.Bill) error
	UpdateTotal(ctx context.Context, billID int64, total float64) error
	WithTx(tx *gorm.DB) BillRepository
}

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{db: db}
}


func (r *billRepository) Create(ctx context.Context, bill *models.Bill) error {
	return r.db.WithContext(ctx).Create(bill).Error
}


func (r *billRepository) UpdateTotal(ctx context.Context,billID int64,total float64,) error {
	return r.db.WithContext(ctx).
		Model(&models.Bill{}).
		Where("id = ?", billID).
		Update("total_amount", total).
		Error
}

func (r *billRepository) WithTx(tx *gorm.DB) BillRepository {
	return &billRepository{db: tx}
}
