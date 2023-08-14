// usecases/payment_usecase_test.go
package usecases

import (
	"promptlabth/ms-payments/entities"
	"testing"
)

type MockPaymentRepository struct{}

func (m *MockPaymentRepository) Store(payment entities.Payment) error {
	return nil // Simulate success
}

func TestProcessPayment(t *testing.T) {
	repo := &MockPaymentRepository{}
	usecase := NewPaymentUsecase(repo) // use the constructor function to create the use case

	stripeID := "stripe_id"

	err := usecase.ProcessPayment(entities.Payment{
		UserID:             1,
		PaymentMethodID:    2,
		Coin:               100.0,
		TransactionStripeID: &stripeID,
		FeatureID:          nil,
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
