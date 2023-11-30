package interfaces

import "github.com/promptlabth/ms-payments/entities"

type PaymentSubscriptionUseCase interface {
	ProcessSubscriptionPayments(payment *entities.PaymentSubscription) error
	GetSubscriptionPaymentBySubscriptionID(payment *entities.PaymentSubscription, subscriptionID string) error
	UpdateSubscriptionPayment(payment *entities.PaymentSubscription) error
}

type PaymentSubscriptionRepository interface {
	Store(payment *entities.PaymentSubscription) error
	Get(payment *entities.PaymentSubscription, paymentIntentId string) error
	UpdateSubscriptionPayment(payment *entities.PaymentSubscription) error
}
