package usecases

import (
	"errors"
	"fmt"
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
)

// type NewPaymentSubscriptionsUsecase interface

type paymentSubscriptionImpl struct {
	Repository interfaces.PaymentSubscriptionRepository
}

func (u *paymentSubscriptionImpl) ProcessSubscriptionPayments(subscriptions_payment entities.PaymentSubscription) error {
	if subscriptions_payment.PaymentIntentId == "" {
		return errors.New("missing Stripe ID")
	}

	fmt.Println("subscriptions_payment")
	fmt.Println(subscriptions_payment)

	return u.Repository.Store(subscriptions_payment)
}

func (u *paymentSubscriptionImpl) GetSubscriptionPaymentByPaymentIntentId(payment *entities.PaymentSubscription, paymentIntentId string) (err error) {
	handleErr := u.Repository.Get(payment, paymentIntentId)
	return handleErr
}

func NewPaymentSubscriptionUsecase(repo interfaces.PaymentSubscriptionRepository) interfaces.PaymentSubscriptionUseCase {
	return &paymentSubscriptionImpl{Repository: repo}
}
