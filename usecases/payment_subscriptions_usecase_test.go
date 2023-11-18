package usecases

import (
	"errors"
	"testing"

	"github.com/promptlabth/ms-payments/entities"
)

type MockPaymentSubscriptionsRepository struct {
	ShouldReturnError bool
}

func (m *MockPaymentSubscriptionsRepository) Store(payment_subscriptions entities.PaymentSubscription) error {
	if m.ShouldReturnError {
		return errors.New("repository error")
	}
	return nil // Simulate success
}

func (m *MockPaymentSubscriptionsRepository) Get(payment *entities.PaymentSubscription, paymentIntentId string) error {
	if m.ShouldReturnError {
		return errors.New("repository error")
	}
	return nil // Simulate success
}

func TestSubscriptionPaymentProcess(t *testing.T) {
	tests := []struct {
		name                string
		repoShouldReturnErr bool
		input               entities.PaymentSubscription
		wantErr             bool
	}{
		{
			name:    "success",
			input:   entities.PaymentSubscription{SubscriptionID: "stripe_id"},
			wantErr: false,
		},
		{
			name:                "repository error",
			repoShouldReturnErr: true,
			input:               entities.PaymentSubscription{SubscriptionID: "stripe_id"},
			wantErr:             true,
		},
		{
			name:    "invalid input - missing stripe_id",
			input:   entities.PaymentSubscription{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockPaymentSubscriptionsRepository{ShouldReturnError: tt.repoShouldReturnErr}
			usecase := NewPaymentSubscriptionUsecase(repo)

			err := usecase.ProcessSubscriptionPayments(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionPaymentProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
