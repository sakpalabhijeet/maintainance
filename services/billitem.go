package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type BillItemService interface {
	Create(ctx context.Context, item *models.BillItem) error
	GetByBillID(ctx context.Context, billID int64) ([]models.BillItem, error)
	DeleteByBillID(ctx context.Context, billID int64) error
}

type billItemService struct {
	repo repo.BillItemRepository
}

func NewBillItemService(repo repo.BillItemRepository,) BillItemService {
	return &billItemService{
		repo: repo,
	}
}


func (s *billItemService) Create(ctx context.Context,item *models.BillItem,) error {
	return s.repo.Create(ctx, item)
}

func (s *billItemService) GetByBillID(ctx context.Context,billID int64,) ([]models.BillItem, error) {
	return s.repo.GetByBillID(ctx, billID)
}


func (s *billItemService) DeleteByBillID(ctx context.Context,billID int64,) error {
	return s.repo.DeleteByBillID(ctx, billID)
}
