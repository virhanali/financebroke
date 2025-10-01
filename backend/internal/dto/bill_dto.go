package dto

import "finance-app/backend/internal/entity"

type BillCreateRequest struct {
	Name         string  `json:"name" binding:"required"`
	Amount       float64 `json:"amount" binding:"required,min=0"`
	DueDate      string  `json:"due_date" binding:"required"`
	Description  string  `json:"description"`
	RemindBefore int     `json:"remind_before"`
}

type BillUpdateRequest struct {
	Name         string  `json:"name"`
	Amount       float64 `json:"amount"`
	DueDate      string  `json:"due_date"`
	Description  string  `json:"description"`
	Status       string  `json:"status"`
	RemindBefore int     `json:"remind_before"`
}

type BillResponse struct {
	entity.Bill
}