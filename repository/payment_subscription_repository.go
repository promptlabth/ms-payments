// payment_subscription_repository.go
package repository

import (
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentSubscriptionsRepository struct {
	conn *gorm.DB
}

type paymentScriptionRepository struct {
	conn *gorm.DB
}

// Store implements interfaces.PaymentSubscriptionRepository.
func (t *PaymentSubscriptionsRepository) Store(payment entities.PaymentSubscription) error {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	newPayment := entities.PaymentSubscription{
		TransactionStripeID: payment.TransactionStripeID,
		PaymentMethod:       payment.PaymentMethod,
		StartDatetime:       now.String(),
		EndDatetime:         oneMonthLater.GoString(),
		Plan:                payment.Plan,
		SubscriptionStatus:  "active",
		User:                payment.User,
		Datetime:            now.GoString(),
	}
	if err := t.conn.Create(newPayment).Error; err != nil {
		return err
	}
	return nil
}

func NewPaymentScriptionRepository(connection *gorm.DB) interfaces.PaymentSubscriptionRepository {
	return &PaymentSubscriptionsRepository{
		conn: connection,
	}
}
