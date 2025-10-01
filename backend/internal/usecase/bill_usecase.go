package usecase

import (
	"financebroke/backend/internal/dto"
	"financebroke/backend/internal/entity"
	"financebroke/backend/internal/repository"
	"time"
)

type BillUsecase interface {
	CreateBill(userID uint, req *dto.BillCreateRequest) (entity.Bill, error)
	GetBill(userID, billID uint) (entity.Bill, error)
	GetUserBills(userID uint) ([]entity.Bill, error)
	GetUpcomingBills(userID uint) ([]entity.Bill, error)
	UpdateBill(userID, billID uint, req *dto.BillUpdateRequest) (entity.Bill, error)
	DeleteBill(userID, billID uint) error
	GetDashboard(userID uint) (*dto.DashboardResponse, error)
}

type billUsecase struct {
	billRepo repository.BillRepository
}

func NewBillUsecase(billRepo repository.BillRepository) BillUsecase {
	return &billUsecase{billRepo: billRepo}
}

func (u *billUsecase) CreateBill(userID uint, req *dto.BillCreateRequest) (entity.Bill, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return entity.Bill{}, err
	}

	remindBefore := 3
	if req.RemindBefore > 0 {
		remindBefore = req.RemindBefore
	}

	bill := entity.Bill{
		UserID:       userID,
		Name:         req.Name,
		Amount:       req.Amount,
		DueDate:      dueDate,
		Description:  req.Description,
		Status:       "unpaid",
		RemindBefore: remindBefore,
	}

	return u.billRepo.Create(bill)
}

func (u *billUsecase) GetBill(userID, billID uint) (entity.Bill, error) {
	return u.billRepo.FindByID(billID, userID)
}

func (u *billUsecase) GetUserBills(userID uint) ([]entity.Bill, error) {
	return u.billRepo.FindByUserID(userID)
}

func (u *billUsecase) GetUpcomingBills(userID uint) ([]entity.Bill, error) {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	return u.billRepo.FindUpcomingBills(userID, now, oneMonthLater)
}

func (u *billUsecase) UpdateBill(userID, billID uint, req *dto.BillUpdateRequest) (entity.Bill, error) {
	bill, err := u.billRepo.FindByID(billID, userID)
	if err != nil {
		return entity.Bill{}, err
	}

	if req.Name != "" {
		bill.Name = req.Name
	}
	if req.Amount > 0 {
		bill.Amount = req.Amount
	}
	if req.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", req.DueDate)
		if err != nil {
			return entity.Bill{}, err
		}
		bill.DueDate = dueDate
	}
	if req.Description != "" {
		bill.Description = req.Description
	}
	if req.Status != "" {
		bill.Status = req.Status
	}
	if req.RemindBefore > 0 {
		bill.RemindBefore = req.RemindBefore
	}

	return u.billRepo.Update(bill)
}

func (u *billUsecase) DeleteBill(userID, billID uint) error {
	return u.billRepo.Delete(billID, userID)
}

func (u *billUsecase) GetDashboard(userID uint) (*dto.DashboardResponse, error) {
	stats, err := u.billRepo.GetDashboardStats(userID)
	if err != nil {
		return nil, err
	}

	upcomingBills, err := u.billRepo.FindUpcomingBills(userID, time.Now(), time.Now().AddDate(0, 1, 0))
	if err != nil {
		return nil, err
	}

	recentBills, err := u.billRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := &dto.DashboardResponse{
		TotalBills:    stats.TotalBills,
		PaidBills:     stats.PaidBills,
		UnpaidBills:   stats.UnpaidBills,
		OverdueBills:  stats.OverdueBills,
		UpcomingBills: upcomingBills,
		RecentBills:   recentBills,
		Summary: dto.BillSummary{
			TotalAmount:    stats.TotalAmount,
			PaidAmount:     stats.PaidAmount,
			UnpaidAmount:   stats.UnpaidAmount,
			OverdueAmount:  stats.OverdueAmount,
		},
	}

	return response, nil
}