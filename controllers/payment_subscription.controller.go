package controllers

import (
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
	"promptlabth/ms-payments/usecases"

	"github.com/gin-gonic/gin"
)

type PaymentSubscriptionController struct {
	paymentSubscriptionUsecase usecases.PaymentSubscriptionUsecase
}


func (p *PaymentSubscriptionController) CreatePaymentSubscription(c *gin.Context) {
	var payment entities.PaymentSubscription
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	if err := p.paymentSubscriptionUsecase.ProcessSubscriptionPayments(payment); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save payment"})
		return
	}
	c.JSON(200, gin.H{"status": "Payment data received and saved"})
}

func NewPaymentScriptionController(usecase interfaces.PaymentSubscriptionUseCase) *PaymentSubscriptionController {
	return &PaymentSubscriptionController{
		paymentSubscriptionUsecase: usecase,
	}
}