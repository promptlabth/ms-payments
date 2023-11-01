// File: stripe_service.go
package services

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func ConfirmPaymentIntent() (bool, error) {
	stripe.Key = os.Getenv("STRIPE_KEY")
	paramsGet := &stripe.PaymentIntentParams{
		// ClientSecret: , //something from client?
	}
	pi, err := paymentintent.Get("pi_3O4N9jAom1IgIvKK1k1LMSgL", paramsGet)
	// To create a PaymentIntent for confirmation, see our guide at:
	// https://stripe.com/docs/payments/payment-intents/creating-payment-intents#creating-for-automatic
	fmt.Print(pi.ClientSecret)
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}

	pi, err = paymentintent.Confirm("pi_3O4N9jAom1IgIvKK1k1LMSgL", params)
	if err != nil {
		return false, err
	}
	if pi.Status == stripe.PaymentIntentStatusSucceeded {
		// Payment was successful
		return true, nil
	}

	// You can add more detailed handling here based on the PaymentIntent's status
	// For example, handle "requires_action" or "requires_payment_method" statuses

	// If the status is not "succeeded", then the payment was not successful
	return false, fmt.Errorf("payment failed with status: %s", pi.Status)
}
