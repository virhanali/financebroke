package entity

import "time"

type Bill struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	Name         string    `json:"name"`
	Amount       float64   `json:"amount"`
	DueDate      time.Time `json:"due_date"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	RemindBefore int       `json:"remind_before"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}