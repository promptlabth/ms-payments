// entities/subscriptions_payments.go
package entities

type PaymentSubscription struct {
	Id                  int `gorm:"primaryKey;autoIncrement:true"`
	TransactionStripeID string
	Datetime            string
	StartDatetime       string
	EndDatetime         string
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

func (b *PaymentSubscription) TableName() string {
	return "subscriptions_payments"
}
