package tests

import (
	"testing"
	"time"

	"github.com/aleksanderpalamar/pixpay/internal/core/payment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPaymentRepository is a mock implementation of the PaymentRepository interface
type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(payment *payment.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) GetByID(id string) (*payment.Payment, error) {
	args := m.Called(id)
	return args.Get(0).(*payment.Payment), args.Error(1)
}

func (m *MockPaymentRepository) Update(payment *payment.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func TestCreatePayment(t *testing.T) {
	mockRepo := new(MockPaymentRepository)
	service := payment.NewService(mockRepo)

	t.Run("success", func(t *testing.T) {
		mockPayment := &payment.Payment{
			ID:        "some-unique-id",
			Amount:    100.0,
			Recipient: "recipient_id",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mockRepo.On("Create", mock.AnythingOfType("*payment.Payment")).Return(nil)

		payment, err := service.CreatePayment(100.0, "recipient_id")

		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, mockPayment.Amount, payment.Amount)
		assert.Equal(t, mockPayment.Recipient, payment.Recipient)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid amount", func(t *testing.T) {
		payment, err := service.CreatePayment(-1, "recipient_id")

		assert.Error(t, err)
		assert.Nil(t, payment)
		assert.Equal(t, "invalid amount", err.Error())
	})

	t.Run("empty recipient", func(t *testing.T) {
		payment, err := service.CreatePayment(100.0, "")

		assert.Error(t, err)
		assert.Nil(t, payment)
		assert.Equal(t, "recipient cannot be empty", err.Error())
	})
}
