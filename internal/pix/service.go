package pix

import (
	"database/sql"
)

type PixPayment struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Receiver    string  `json:"receiver"`
}

type PixService interface {
	CreatePayment(payment PixPayment) error
}

type pixService struct {
	db *sql.DB
}

func NewPixService(db *sql.DB) PixService {
	return &pixService{db: db}
}

func (s *pixService) CreatePayment(payment PixPayment) error {
	// Implementação para salvar o pagamento no banco de dados
	_, err := s.db.Exec("INSERT INTO payments (id, amount, description, receiver) VALUES ($1, $2, $3, $4)",
		payment.ID, payment.Amount, payment.Description, payment.Receiver)
	return err
}
