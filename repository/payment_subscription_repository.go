// payment_subscription_repository.go
package repository

import (
	"time"

	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"

	"gorm.io/gorm"
)

type PaymentSubscriptionsRepository struct {
	conn *gorm.DB
}

// Store implements interfaces.PaymentSubscriptionRepository.
func (t *PaymentSubscriptionsRepository) Store(payment entities.PaymentSubscription) error {
	now := time.Now()
	oneMonthLater := now.AddDate(0, 1, 0)
	newPayment := entities.PaymentSubscription{
		SubscriptionID:     payment.SubscriptionID,
		PaymentMethodID:    payment.PaymentMethodID,
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

func (t *PaymentSubscriptionsRepository) Get(payment *entities.PaymentSubscription, subscriptionID string) error {
	if err := t.conn.Where("subscription_id = ?", subscriptionID).First(&payment).Error; err != nil {
		return err
	}
	return nil
}

func NewPaymentScriptionRepository(connection *gorm.DB) interfaces.PaymentSubscriptionRepository {
	return &PaymentSubscriptionsRepository{
		conn: connection,
	}
}
