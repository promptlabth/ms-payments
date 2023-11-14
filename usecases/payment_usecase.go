// usecases/payment_usecase.go
package usecases

import "github.com/promptlabth/ms-payments/entities"

type PaymentUsecase interface {
	ProcessPayment(payment entities.Payment) error
}

type paymentUsecaseImpl struct {
	Repository PaymentRepository
}

func (u *paymentUsecaseImpl) ProcessPayment(payment entities.Payment) error {
	return u.Repository.Store(payment)
}

func NewPaymentUsecase(repo PaymentRepository) PaymentUsecase {
	return &paymentUsecaseImpl{Repository: repo}
}

type PaymentRepository interface {
	Store(payment entities.Payment) error
}
