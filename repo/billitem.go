package repo

import (
	"Maintainance/models"
	"context"
	"gorm.io/gorm"
)

type BillItemRepository interface {
	Create(ctx context.Context, item *models.BillItem) error
	GetByBillID(ctx context.Context, billID int64) ([]models.BillItem, error)
	DeleteByBillID(ctx context.Context, billID int64) error
	WithTx(tx *gorm.DB) BillItemRepository
}

type billItemRepository struct {
	db *gorm.DB
}

func NewBillItemRepository(db *gorm.DB) BillItemRepository {
	return &billItemRepository{db: db}
}

func (r *billItemRepository) Create(ctx context.Context,item *models.BillItem,) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *billItemRepository) GetByBillID(ctx context.Context,billID int64,) ([]models.BillItem, error) {

	var items []models.BillItem
	err := r.db.WithContext(ctx).
		Where("bill_id = ?", billID).
		Find(&items).Error

	return items, err
}

func (r *billItemRepository) DeleteByBillID(ctx context.Context,billID int64,) error {
	return r.db.WithContext(ctx).
		Where("bill_id = ?", billID).
		Delete(&models.BillItem{}).
		Error
}

func (r *billItemRepository) WithTx(tx *gorm.DB) BillItemRepository {
	return &billItemRepository{db: tx}
}

