package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
)

type BillService interface {
	CreateBill(ctx context.Context, bill *models.Bill)error
}

type billService struct{
	repo repo.BillRepository
}

func NewBillService( repo repo.BillRepository) BillService{
	return &billService{repo: repo}
}

func (s *billService)CreateBill(ctx context.Context, bill *models.Bill)error{
	return s.repo.Create(ctx, bill)
}