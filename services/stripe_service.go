// File: stripe_service.go
package services

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func ConfirmPaymentIntent(paymentIntentID, paymentMethodID string) (bool, error) {
    // Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")

    // Retrieve the PaymentIntent to check its status before confirming
	paramsGet := &stripe.PaymentIntentParams{}
	pi, err := paymentintent.Get(paymentIntentID, paramsGet)
	if err != nil {
		return false, fmt.Errorf("error retrieving payment intent: %v", err)
	}
    
    // Check if the PaymentIntent is already succeeded
    if pi.Status == stripe.PaymentIntentStatusSucceeded {
        return true, nil
    }

    // Confirm the PaymentIntent if it is not succeeded
	paramsConfirm := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(paymentMethodID),
	}
	pi, err = paymentintent.Confirm(paymentIntentID, paramsConfirm)
	if err != nil {
		return false, fmt.Errorf("error confirming payment intent: %v", err)
	}
	
    // Check the status of the PaymentIntent after confirmation
	if pi.Status == stripe.PaymentIntentStatusSucceeded {
		return true, nil
	} else {
		return false, fmt.Errorf("payment failed with status: %s", pi.Status)
	}
}

