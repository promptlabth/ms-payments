// File: stripe_service.go
package services

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/invoice"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/subscription"
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
	} else {
		return false, fmt.Errorf("payment failed with status: %s", pi.Status)
	}
}

func CreateCustomer(email string, name string, firebaseId string) (*stripe.Customer, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")

	// customer Params
	params := &stripe.CustomerParams{
		Email:       &email,
		Name:        &name,
		Description: stripe.String(firebaseId),
	}

	c, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return c, nil

}

func CreateCheckoutSession(
	prize string,
	mode string,
	paymentMethodType []string,
	customerStripeId string,
	originWeb string,
	plan int,
) (*stripe.CheckoutSession, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")

	// encrypt a plan code
	ciperPlan, err := GetAESEncrypted(strconv.Itoa(plan))
	if err != nil {
		return nil, err
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(prize),
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(mode),
		PaymentMethodTypes: stripe.StringSlice(
			paymentMethodType,
		),
		Customer:   &customerStripeId,
		SuccessURL: stripe.String(fmt.Sprintf("%s/subscription/success?session_id={CHECKOUT_SESSION_ID}&plan=%s", originWeb, ciperPlan)),
		CancelURL:  stripe.String(fmt.Sprintf("%s/cancel?cancle={cancle}", originWeb)),
	}

	s, err := session.New(params)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func CheckCheckoutSessionId(sessionId string) (*stripe.CheckoutSession, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")

	s, err := session.Get(
		sessionId, nil,
	)
	if err != nil {
		return nil, err
	}
	return s, nil

}

func GetPriceBySubscriptionID(subscriptionID string) (*stripe.SubscriptionItem, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")

	s, err := subscription.Get(
		subscriptionID,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return s.Items.Data[0], nil
}

func CancelSubscriptionBySubID(subscriptionID string) (*stripe.Subscription, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")
	s, _ := subscription.Cancel(
		subscriptionID,
		nil,
	)
	return s, nil
}

func ListSubscriptionByCustomerID(customerID string) (*[]stripe.Subscription, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")
	params := &stripe.SubscriptionListParams{
		Customer: &customerID,
	}
	i := subscription.List(params)
	var subscriptionList []stripe.Subscription
	for i.Next() {
		subscriptionList = append(subscriptionList, *i.Subscription())
	}
	return &subscriptionList, nil
}

func ListInvoicesByCustomerID(customerID string) (*[]stripe.Invoice, error) {
	// Set your Stripe secret API key
	stripe.Key = os.Getenv("STRIPE_KEY")
	params := &stripe.InvoiceListParams{
		Customer: &customerID,
	}
	i := invoice.List(params)

	var invoices []stripe.Invoice
	for i.Next() {
		invoices = append(invoices, *i.Invoice())
	}
	return &invoices, nil
}
