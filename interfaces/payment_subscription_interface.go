package interfaces

import "promptlabth/ms-payments/entities"

type PaymentSubscriptionUseCase interface {
	ProcessSubscriptionPayments(payment entities.PaymentSubscription) error
}

type PaymentSubscriptionRepository interface {
	Store(payment entities.PaymentSubscription) error
}