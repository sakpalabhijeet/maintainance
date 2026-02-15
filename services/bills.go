package services

import (
	"Maintainance/models"
	"Maintainance/repo"
	"context"
	"time"
)

type BillService interface {
    GenerateMonthlyBills(ctx context.Context, billMonth time.Time, dueDate time.Time) error
}

type billService struct {
    billRepo repo.BillRepository
    flatRepo repo.FlatRepository
    rateRepo repo.MaintenanceRateRepository
    itemRepo repo.BillItemRepository
}


func NewBillService(
    billRepo repo.BillRepository,
    flatRepo repo.FlatRepository,
    rateRepo repo.MaintenanceRateRepository,
    itemRepo repo.BillItemRepository,
) BillService {
    return &billService{
        billRepo: billRepo,
        flatRepo: flatRepo,
        rateRepo: rateRepo,
        itemRepo: itemRepo,
    }
}

func (s *billService) GenerateMonthlyBills(
	ctx context.Context,
	billMonth time.Time,
	dueDate time.Time,
) error {

	flats, err := s.flatRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	rates, err := s.rateRepo.GetActiveRates(ctx, billMonth)
	if err != nil {
		return err
	}

	for _, flat := range flats {

		bill := models.Bill{
			FlatID:    flat.Id,
			BillMonth: billMonth,
			Status:    "pending",
			DueDate:   dueDate,
		}

		if err := s.billRepo.Create(ctx, &bill); err != nil {
			return err
		}

		total := 0.0

		for _, rate := range rates {

			baseAmount := flat.AreaSqFt * rate.RatePerSqFt
			gstAmount := baseAmount * (rate.GSTPercent / 100)
			finalAmount := baseAmount + gstAmount

			total += finalAmount

			item := models.BillItem{
				BillID:      bill.BillID,
				Description: rate.RateType,
				Rate:        rate.RatePerSqFt,
				Quantity:    flat.AreaSqFt,
				Amount:      finalAmount,
			}

			if err := s.itemRepo.Create(ctx, &item); err != nil {
				return err
			}
		}

		if err := s.billRepo.UpdateTotal(ctx, bill.BillID, total); err != nil {
			return err
		}
	}

	return nil
}

func (s *billService) ApplyLatePenalty(
	ctx context.Context,
	bill models.Bill,
	penaltyPercent float64,
) error {

	if time.Now().After(bill.DueDate) && bill.Status == "pending" {

		penaltyAmount := bill.TotalAmount * (penaltyPercent / 100)

		item := models.BillItem{
			BillID:      bill.BillID,
			Description: "Late Payment Penalty",
			Rate:        penaltyPercent,
			Quantity:    1,
			Amount:      penaltyAmount,
		}

		if err := s.itemRepo.Create(ctx, &item); err != nil {
			return err
		}

		newTotal := bill.TotalAmount + penaltyAmount
		return s.billRepo.UpdateTotal(ctx, bill.BillID, newTotal)
	}

	return nil
}

