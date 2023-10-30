package usecases

import (
	"errors"
	"promptlabth/ms-payments/entities"
)

// type NewPaymentSubscriptionsUsecase interface

type PaymentSubscriptionsRepository interface {
	Store(subscriptionPayment entities.PaymentSubscription) error
}

type PaymentSubscriptionUsecase interface {
	ProcessSubscriptionPayments(payment entities.PaymentSubscription) error
}


type paymentSubscriptionImpl struct {
	Repository PaymentSubscriptionsRepository
}


func (u *paymentSubscriptionImpl) ProcessSubscriptionPayments(subscriptions_payment entities.PaymentSubscription) error {
	if subscriptions_payment.TransactionStripeID == "" {
		return errors.New("missing Stripe ID")
	}

	return u.Repository.Store(subscriptions_payment)
}

func NewPaymentSubscriptionUsecase(repo PaymentSubscriptionsRepository) PaymentSubscriptionUsecase {
	return &paymentSubscriptionImpl{Repository: repo}
}