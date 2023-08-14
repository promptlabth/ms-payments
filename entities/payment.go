// entities/payment.go
package entities

type Payment struct {
	TransactionID       int
	UserID              int
	PaymentMethodID     int
	Coin                float64
	TransactionStripeID *string
	Datetime            string
	FeatureID           *int
}
