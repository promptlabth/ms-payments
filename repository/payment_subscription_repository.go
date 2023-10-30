// payment_subscription_repository.go
package repository

import (
	"database/sql"
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentSubscriptionsRepository struct {
	DB *sql.DB
}

type paymentScriptionRepository struct {
	conn *gorm.DB
}

// Store implements interfaces.PaymentSubscriptionRepository.
func (*paymentScriptionRepository) Store(payment entities.PaymentSubscription) error {
	panic("unimplemented")
}

func (r *PaymentSubscriptionsRepository) Store(payment entities.PaymentSubscription) error {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	_, err := r.DB.Exec(`INSERT INTO subscriptions_payments (UserID, PaymentMethodsID, PlanID, TransactionStripeID, Datetime, StartDatetime, EndDatetime, SubscriptionStatus) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		payment.UserID, payment.PaymentMethodID, payment.Plan.Id, payment.TransactionStripeID, now, now, oneMonthLater, "active")
	return err
}

func NewPaymentScriptionRepository(connection *gorm.DB) interfaces.PaymentSubscriptionRepository {
	return &paymentScriptionRepository{
		conn: connection,
	}
}
