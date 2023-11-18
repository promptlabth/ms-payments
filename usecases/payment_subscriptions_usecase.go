package usecases

import (
	"errors"
	"fmt"

	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
)

// type NewPaymentSubscriptionsUsecase interface

type paymentSubscriptionImpl struct {
	Repository interfaces.PaymentSubscriptionRepository
}

func (u *paymentSubscriptionImpl) ProcessSubscriptionPayments(subscriptions_payment entities.PaymentSubscription) error {
	if subscriptions_payment.SubscriptionID == "" {
		return errors.New("missing Stripe ID")
	}

	fmt.Println("subscriptions_payment")
	fmt.Println(subscriptions_payment)

	return u.Repository.Store(subscriptions_payment)
}

func (u *paymentSubscriptionImpl) GetSubscriptionPaymentBySubscriptionID(payment *entities.PaymentSubscription, subscriptionID string) (err error) {
	handleErr := u.Repository.Get(payment, subscriptionID)
	return handleErr
}

func NewPaymentSubscriptionUsecase(repo interfaces.PaymentSubscriptionRepository) interfaces.PaymentSubscriptionUseCase {
	return &paymentSubscriptionImpl{Repository: repo}
}
