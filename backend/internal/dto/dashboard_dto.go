package dto

import "financebroke/backend/internal/entity"

type DashboardResponse struct {
	TotalBills    int64          `json:"total_bills"`
	PaidBills     int64          `json:"paid_bills"`
	UnpaidBills   int64          `json:"unpaid_bills"`
	OverdueBills  int64          `json:"overdue_bills"`
	UpcomingBills []entity.Bill  `json:"upcoming_bills"`
	RecentBills   []entity.Bill  `json:"recent_bills"`
	Summary       BillSummary    `json:"summary"`
}

type BillSummary struct {
	TotalAmount    float64 `json:"total_amount"`
	PaidAmount     float64 `json:"paid_amount"`
	UnpaidAmount   float64 `json:"unpaid_amount"`
	OverdueAmount  float64 `json:"overdue_amount"`
}