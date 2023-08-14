// usecases/payment_usecase.go
package usecases

import (
	"promptlabth/ms-payments/entities"
)

type PaymentRepository interface {
	Store(payment entities.Payment) error
}

type PaymentUsecase struct {
	Repository PaymentRepository
}

func (u *PaymentUsecase) ProcessPayment(payment entities.Payment) error {
	return u.Repository.Store(payment)
}
