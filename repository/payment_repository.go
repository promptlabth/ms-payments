// gateways/payment_repository.go
package repository

import (
	"database/sql"
	"promptlabth/ms-payments/entities"
	"time"
)

type PaymentRepository struct {
	DB *sql.DB
}

func (r *PaymentRepository) Store(payment entities.Payment) error {
	_, err := r.DB.Exec(`INSERT INTO payments (UserID, PaymentMethodsID, Coin, Transection_Stripe_Id, Datetime, FeatureID) VALUES ($1, $2, $3, $4, $5, $6)`,
		payment.UserID, payment.PaymentMethodID, payment.Coin, payment.TransactionStripeID, time.Now(), payment.FeatureID)
	return err
}
