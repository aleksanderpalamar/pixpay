package payment

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Recipient string    `json:"recipient"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentRepository interface {
	Create(payment *Payment) error
	GetByID(is string) (*Payment, error)
	Update(payment *Payment) error
}

type Service struct {
	repository PaymentRepository
}

func NewService(repo PaymentRepository) *Service {
	return &Service{
		repository: repo,
	}
}

// Create creates a new payment
func (s *Service) CreatePayment(amount float64, recipient string) (*Payment, error) {
	if amount <= 0 {
		return nil, errors.New("invalid amount")
	}
	if recipient == "" {
		return nil, errors.New("recipient cannot be empty")
	}
	payment := &Payment{
		ID:        generateID(),
		Amount:    amount,
		Recipient: recipient,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repository.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *Service) GetPaymentByID(id string) (*Payment, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	payment, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *Service) UpdatePayment(payment *Payment) error {
	if payment.ID == "" {
		return errors.New("id cannot be empty")
	}
	if payment.Amount <= 0 {
		return errors.New("invalid amount")
	}
	if payment.Recipient == "" {
		return errors.New("recipient cannot be empty")
	}

	payment.UpdatedAt = time.Now()

	err := s.repository.Update(payment)
	if err != nil {
		return err
	}

	return nil
}

// Helper function to generate a unique ID (for simplicity, using a UUID library)
func generateID() string {
	return uuid.New().String()
}
