package repository

import (
	"database/sql"
	"finance-app/backend/internal/entity"
	"time"
)

type BillRepository interface {
	Create(bill entity.Bill) (entity.Bill, error)
	FindByID(id, userID uint) (entity.Bill, error)
	FindByUserID(userID uint) ([]entity.Bill, error)
	FindUpcomingBills(userID uint, startDate, endDate time.Time) ([]entity.Bill, error)
	Update(bill entity.Bill) (entity.Bill, error)
	Delete(id, userID uint) error
	GetDashboardStats(userID uint) (DashboardStats, error)
}

type DashboardStats struct {
	TotalBills    int64
	PaidBills     int64
	UnpaidBills   int64
	OverdueBills  int64
	TotalAmount   float64
	PaidAmount    float64
	UnpaidAmount  float64
	OverdueAmount float64
}

type billRepository struct {
	db *sql.DB
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}

func (r *billRepository) Create(bill entity.Bill) (entity.Bill, error) {
	query := `
		INSERT INTO bills (user_id, name, amount, due_date, description, status, remind_before)
		VALUES ($1, $2, $3, $4, $5, 'unpaid', $6)
		RETURNING id, user_id, name, amount, due_date, description, status, remind_before, created_at, updated_at
	`

	err := r.db.QueryRow(query, bill.UserID, bill.Name, bill.Amount, bill.DueDate, bill.Description, bill.RemindBefore).Scan(
		&bill.ID, &bill.UserID, &bill.Name, &bill.Amount, &bill.DueDate, &bill.Description,
		&bill.Status, &bill.RemindBefore, &bill.CreatedAt, &bill.UpdatedAt,
	)

	if err != nil {
		return entity.Bill{}, err
	}

	return bill, nil
}

func (r *billRepository) FindByID(id, userID uint) (entity.Bill, error) {
	query := `
		SELECT id, user_id, name, amount, due_date, description, status, remind_before, created_at, updated_at
		FROM bills
		WHERE id = $1 AND user_id = $2
	`

	var bill entity.Bill
	var description sql.NullString
	err := r.db.QueryRow(query, id, userID).Scan(
		&bill.ID, &bill.UserID, &bill.Name, &bill.Amount, &bill.DueDate, &description,
		&bill.Status, &bill.RemindBefore, &bill.CreatedAt, &bill.UpdatedAt,
	)

	if err != nil {
		return entity.Bill{}, err
	}

	if description.Valid {
		bill.Description = description.String
	}

	return bill, nil
}

func (r *billRepository) FindByUserID(userID uint) ([]entity.Bill, error) {
	query := `
		SELECT id, user_id, name, amount, due_date, description, status, remind_before, created_at, updated_at
		FROM bills
		WHERE user_id = $1
		ORDER BY due_date ASC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bills []entity.Bill
	for rows.Next() {
		var bill entity.Bill
		var description sql.NullString
		err := rows.Scan(
			&bill.ID, &bill.UserID, &bill.Name, &bill.Amount, &bill.DueDate, &description,
			&bill.Status, &bill.RemindBefore, &bill.CreatedAt, &bill.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			bill.Description = description.String
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

func (r *billRepository) FindUpcomingBills(userID uint, startDate, endDate time.Time) ([]entity.Bill, error) {
	query := `
		SELECT id, user_id, name, amount, due_date, description, status, remind_before, created_at, updated_at
		FROM bills
		WHERE user_id = $1 AND due_date BETWEEN $2 AND $3 AND status != 'paid'
		ORDER BY due_date ASC
	`

	rows, err := r.db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bills []entity.Bill
	for rows.Next() {
		var bill entity.Bill
		var description sql.NullString
		err := rows.Scan(
			&bill.ID, &bill.UserID, &bill.Name, &bill.Amount, &bill.DueDate, &description,
			&bill.Status, &bill.RemindBefore, &bill.CreatedAt, &bill.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if description.Valid {
			bill.Description = description.String
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

func (r *billRepository) Update(bill entity.Bill) (entity.Bill, error) {
	query := `
		UPDATE bills
		SET name = $2, amount = $3, due_date = $4, description = $5, status = $6, remind_before = $7, updated_at = CURRENT_TIMESTAMP
		WHERE id = $1 AND user_id = $8
		RETURNING id, user_id, name, amount, due_date, description, status, remind_before, created_at, updated_at
	`

	var description sql.NullString
	if bill.Description != "" {
		description.String = bill.Description
		description.Valid = true
	}

	err := r.db.QueryRow(query, bill.ID, bill.Name, bill.Amount, bill.DueDate, description, bill.Status, bill.RemindBefore, bill.UserID).Scan(
		&bill.ID, &bill.UserID, &bill.Name, &bill.Amount, &bill.DueDate, &bill.Description,
		&bill.Status, &bill.RemindBefore, &bill.CreatedAt, &bill.UpdatedAt,
	)

	if err != nil {
		return entity.Bill{}, err
	}

	return bill, nil
}

func (r *billRepository) Delete(id, userID uint) error {
	query := `DELETE FROM bills WHERE id = $1 AND user_id = $2`
	result, err := r.db.Exec(query, id, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *billRepository) GetDashboardStats(userID uint) (DashboardStats, error) {
	query := `
		SELECT
			COUNT(*) as total_bills,
			COUNT(CASE WHEN status = 'paid' THEN 1 END) as paid_bills,
			COUNT(CASE WHEN status = 'unpaid' THEN 1 END) as unpaid_bills,
			COUNT(CASE WHEN status = 'overdue' THEN 1 END) as overdue_bills,
			COALESCE(SUM(amount), 0) as total_amount,
			COALESCE(SUM(CASE WHEN status = 'paid' THEN amount END), 0) as paid_amount,
			COALESCE(SUM(CASE WHEN status = 'unpaid' THEN amount END), 0) as unpaid_amount,
			COALESCE(SUM(CASE WHEN status = 'overdue' THEN amount END), 0) as overdue_amount
		FROM bills
		WHERE user_id = $1
	`

	var stats DashboardStats
	err := r.db.QueryRow(query, userID).Scan(
		&stats.TotalBills, &stats.PaidBills, &stats.UnpaidBills, &stats.OverdueBills,
		&stats.TotalAmount, &stats.PaidAmount, &stats.UnpaidAmount, &stats.OverdueAmount,
	)

	if err != nil {
		return DashboardStats{}, err
	}

	return stats, nil
}