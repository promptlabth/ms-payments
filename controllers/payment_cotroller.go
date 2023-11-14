// controllers/payment_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/usecases"
)

type PaymentController struct {
	Usecase usecases.PaymentUsecase
}

func (p *PaymentController) CreatePayment(c *gin.Context) {
	var payment entities.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	if err := p.Usecase.ProcessPayment(payment); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save payment"})
		return
	}
	c.JSON(200, gin.H{"status": "Payment data received and saved"})
}
