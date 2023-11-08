package interfaces

import "promptlabth/ms-payments/entities"

type PaymentSubscriptionUseCase interface {
	ProcessSubscriptionPayments(payment entities.PaymentSubscription) error
	GetSubscriptionPaymentByPaymentIntentId(payment *entities.PaymentSubscription, paymentIntentId string) error
}

type PaymentSubscriptionRepository interface {
	Store(payment entities.PaymentSubscription) error
	Get(payment *entities.PaymentSubscription, paymentIntentId string) error
}
