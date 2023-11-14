package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/promptlabth/ms-payments/entities"

	"github.com/gin-gonic/gin"
)

type mockPaymentSubscriptionUsecase struct{}

func (m *mockPaymentSubscriptionUsecase) ProcessSubscriptionPayments(payment entities.PaymentSubscription) error {
	return nil
}

func (m *mockPaymentSubscriptionUsecase) GetSubscriptionPaymentByPaymentIntentId(payment *entities.PaymentSubscription, paymentIntentId string) error {
	return nil
}

func TestCreatePaymentSubscription(t *testing.T) {
	tests := []struct {
		name       string
		input      entities.PaymentSubscription
		wantStatus int
	}{
		{
			name:       "success",
			input:      entities.PaymentSubscription{},
			wantStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert the input to JSON
			jsonValue, _ := json.Marshal(tt.input)

			// Create a new HTTP request with the JSON body
			req, _ := http.NewRequest("POST", "/some-endpoint", bytes.NewBuffer(jsonValue))

			// Create a new Gin context and assign the HTTP request to it
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request = req

			// Create the controller and call the function
			controller := &PaymentSubscriptionController{paymentSubscriptionUsecase: &mockPaymentSubscriptionUsecase{}}
			controller.CreatePaymentSubscription(c)

			// Assert the response status code
			if c.Writer.Status() != tt.wantStatus {
				t.Errorf("CreatePaymentSubscription() = %v, want %v", c.Writer.Status(), tt.wantStatus)
			}
		})
	}
}
