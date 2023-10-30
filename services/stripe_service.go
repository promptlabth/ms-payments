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

	// To create a PaymentIntent for confirmation, see our guide at:
	// https://stripe.com/docs/payments/payment-intents/creating-payment-intents#creating-for-automatic
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}

	pi, err := paymentintent.Confirm("pi_1FOhymCiQSLSNSsfhI8BdrrR", params)
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


