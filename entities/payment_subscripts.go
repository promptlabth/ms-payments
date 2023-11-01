// entities/subscriptions_payments.go
package entities

import "time"

type PaymentSubscription struct {
	Id                  int `gorm:"primaryKey;autoIncrement:true"`
	TransactionStripeID string
	Datetime            time.Time
	StartDatetime       time.Time
	EndDatetime         time.Time
	SubscriptionStatus  string

	// user
	UserID *uint `valid:"-"`
	User   User  `gorm:"references:Id" valid:"-"`

	// payment
	PaymentMethodID *uint         `valid:"-"`
	PaymentMethod   PaymentMethod `gorm:"references:Id" valid:"-"`

	// plan
	PlanID *uint `valid:"-"`
	Plan   Plan  `gorm:"references:Id" valid:"-"`
}

type PaymentSubscriptionRequest struct {
	TransactionStripeID string

	// user
	UserID *uint `valid:"-"`
	User   User  `gorm:"references:Id" valid:"-"`

	// payment
	PaymentMethodID *uint         `valid:"-"`
	PaymentMethod   PaymentMethod `gorm:"references:Id" valid:"-"`

	// plan
	PlanID *uint `valid:"-"`
	Plan   Plan  `gorm:"references:Id" valid:"-"`
}

func (b *PaymentSubscription) TableName() string {
	return "subscriptions_payments"
}
