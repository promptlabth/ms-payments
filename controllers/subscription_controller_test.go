package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/entities"
)

type mockUserUsecase struct{}

func (m *mockUserUsecase) GetAUserByFirebaseId(out *entities.User, firebaseId string) (err error) {
	return nil
}
func (m *mockUserUsecase) UpdateAUser(out *entities.User, id string) (err error) {
	return nil
}
func (m *mockUserUsecase) GetAUserByStripeID(out *entities.User, stripeId string) error {
	return nil
}

type mockPaymentSubscriptionUsecase struct{}

func (m *mockPaymentSubscriptionUsecase) ProcessSubscriptionPayments(payment *entities.PaymentSubscription) error {
	return nil
}
func (m *mockPaymentSubscriptionUsecase) GetSubscriptionPaymentBySubscriptionID(payment *entities.PaymentSubscription, subscriptionID string) error {
	return nil
}
func (m *mockPaymentSubscriptionUsecase) UpdateSubscriptionPayment(payment *entities.PaymentSubscription) error {
	return nil
}

type mockPlanUsecase struct{}

func (m *mockPlanUsecase) GetAPlan(plan *entities.Plan, id int) error {
	return nil
}
func (m *mockPlanUsecase) GetAPlanByPriceID(plan *entities.Plan, id string) error {
	return nil
}
func (m *mockPlanUsecase) CreateAPlan(plan *entities.Plan) error {
	return nil
}
func (m *mockPlanUsecase) GetAPlanByProdID(plan *entities.Plan, id string) error {
	return nil
}

func TestSuccessfulSubscription(t *testing.T) {
	tests := []struct {
		name       string
		input      entities.SubscriptionReqUrl
		wantStatus int
	}{
		{
			name:       "success",
			input:      entities.SubscriptionReqUrl{},
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
			controller := &SubscriptionReqUrlController{
				userUsecase:                &mockUserUsecase{},
				planUsecase:                &mockPlanUsecase{},
				paymentSubscriptionUsecase: &mockPaymentSubscriptionUsecase{},
			}
			controller.SaveSubscription(c)

			// Assert the response status code
			if c.Writer.Status() != tt.wantStatus {
				t.Errorf("CreatePaymentSubscription() = %v, want %v", c.Writer.Status(), tt.wantStatus)
			}
		})
	}
}
