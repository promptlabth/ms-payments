// entities/payment.go
package entities

type Payment struct {
	Id                  int `gorm:"primaryKey;autoIncrement:true"`
	Coin                int
	TransactionStripeID string
	Datetime            string

	// user
	UserID *uint `valid:"-"`
	User   User  `gorm:"references:Id" valid:"-"`
	// payment
	PaymentMethodID *uint         `valid:"-"`
	PaymentMethod   PaymentMethod `gorm:"references:Id" valid:"-"`
	// feature
	FeatureID *uint   `valid:"-"`
	Feature   Feature `gorm:"references:Id" valid:"-"`
}

func (b *Payment) TableName() string {
	return "payment"
}
