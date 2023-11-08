package controllers

import (
	"fmt"
	"net/http"
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
	"promptlabth/ms-payments/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentSubscriptionController struct {
	paymentSubscriptionUsecase interfaces.PaymentSubscriptionUseCase
}

type PaymentSubscriptionRequestWrapper struct {
	*entities.PaymentSubscriptionRequest
}

func (p *PaymentSubscriptionController) CreatePaymentSubscription(c *gin.Context) {
	var paymentSubscriptionRequest PaymentSubscriptionRequestWrapper

	// Bind the incoming JSON payload to the PaymentSubscriptionRequestWrapper struct
	if err := c.ShouldBindJSON(&paymentSubscriptionRequest); err != nil {
		// Log the error for debugging
		fmt.Println("Error binding JSON:", err)

		// Respond with a 400 Bad Request status code and detailed error message
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": err.Error()})
		return
	}

	// Convert the paymentSubscriptionRequest to PaymentSubscription
	paymentSubscription := paymentSubscriptionRequest.ToPaymentSubscription()

	paymentMethodIDStr := strconv.FormatUint(uint64(*paymentSubscription.PaymentMethodID), 10)

	// Confirm the payment intent
	success, err := services.ConfirmPaymentIntent(paymentSubscriptionRequest.PaymentIntentId, paymentMethodIDStr)
	if err != nil {
		// Log the error for debugging
		fmt.Println("Error confirming payment intent:", err)

		// Respond with a 500 Internal Server Error status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": "Failed to confirm payment intent"})
		return
	}

	if !success {
		// Log the error for debugging
		fmt.Println("Payment failed.")

		// Respond with a 400 Bad Request status code
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": "Payment was not successful"})
		return
	}

	// validate a PaymentIntentId (find a payment Intent ID)
	var payment entities.PaymentSubscription
	if err := p.paymentSubscriptionUsecase.GetSubscriptionPaymentByPaymentIntentId(&payment, paymentSubscriptionRequest.PaymentIntentId); err == nil {
		fmt.Println("Error processing subscription payments:", err)

		// Respond with a 400 Bad Request status code
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "details": "Failed to process. found a duplicate data"})
		return
	}

	// Log the successful payment
	fmt.Println("Payment was successful!")

	// Process the subscription payments
	if err := p.paymentSubscriptionUsecase.ProcessSubscriptionPayments(paymentSubscription); err != nil {
		// Log the error for debugging
		fmt.Println("Error processing subscription payments:", err)

		// Respond with a 500 Internal Server Error status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": "Failed to process subscription payments"})
		return
	}

	// Respond with a 200 OK status code and success message
	c.JSON(http.StatusOK, gin.H{"status": "Payment data received and saved"})
}

func NewPaymentScriptionController(usecase interfaces.PaymentSubscriptionUseCase) *PaymentSubscriptionController {
	return &PaymentSubscriptionController{
		paymentSubscriptionUsecase: usecase,
	}
}

func (r *PaymentSubscriptionRequestWrapper) ToPaymentSubscription() entities.PaymentSubscription {
	now := time.Now()
	paymentMethodID := uint(1)
	return entities.PaymentSubscription{
		PaymentIntentId:    r.PaymentIntentId,
		Datetime:           now,
		StartDatetime:      now,
		EndDatetime:        now.Add(30 * 24 * time.Hour), // For example, 30 days later
		SubscriptionStatus: "active",
		UserID:             r.UserID,
		PaymentMethodID:    &paymentMethodID,
		PlanID:             r.PlanID,
	}
}
