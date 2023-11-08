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

// Store implements interfaces.PaymentSubscriptionRepository.
func (t *PaymentSubscriptionsRepository) Store(payment entities.PaymentSubscription) error {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	newPayment := entities.PaymentSubscription{
		PaymentIntentId:    payment.PaymentIntentId,
		PaymentMethod:      payment.PaymentMethod,
		StartDatetime:      now,
		EndDatetime:        oneMonthLater,
		PlanID:             payment.PlanID,
		SubscriptionStatus: "active",
		UserID:             payment.UserID,
		Datetime:           now,
	}
	if err := t.conn.Create(&newPayment).Error; err != nil {
		return err
	}
	return nil
}

func (t *PaymentSubscriptionsRepository) Get(payment *entities.PaymentSubscription, paymentIntentId string) error {
	if err := t.conn.Raw("SELECT * FROM subscriptions_payments WHERE payment_intent_id = ?", paymentIntentId).Find(&payment).Error; err != nil {
		return err
	}
	return nil
}

func NewPaymentScriptionRepository(connection *gorm.DB) interfaces.PaymentSubscriptionRepository {
	return &PaymentSubscriptionsRepository{
		conn: connection,
	}
}
