package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

type Payment struct {
	ID         int
	SenderID   string
	ReceiverID string
	Amount     float64
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepositoty(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment Payment) error {
	query := `
		INSERT INTO payments (sender_id, receiver_id, amount, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(query, payment.SenderID, payment.ReceiverID, payment.Amount, payment.Status).Scan(&payment.ID, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating payment: %w", err)
	}
	return nil
}

func (r *PaymentRepository) GetPayment(id int) (*Payment, error) {
	query := `
		SELECT id, sender_id, receiver_id, amount, status, created_at, updated_at
		FROM payments
		WHERE id = $1
	`
	var payment Payment
	err := r.db.QueryRow(query, id).Scan(&payment.ID, &payment.ReceiverID, &payment.SenderID, &payment.Amount, &payment.Status, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not execute select query: %v", err)
	}
	return &payment, nil
}

func (r *PaymentRepository) UpdatePayment(payment Payment) error {
	query := `
		UPDATE payments
		SET receiver_id = $2, sender_id = $3, amount = $4, status = $5
		WHERE id = $1
	`
	_, err := r.db.Exec(query, payment.ID, payment.ReceiverID, payment.SenderID, payment.Amount, payment.Status)
	if err != nil {
		return fmt.Errorf("could not execute update query: %v", err)
	}
	return nil
}

func (r *PaymentRepository) DeletePayment(id int) error {
	query := `
		DELETE FROM payments
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not execute delete query: %v", err)
	}
	return nil
}
