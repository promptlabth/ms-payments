// _test/controllers/payment_controller_test.go
package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"promptlabth/ms-payments/entities"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockPaymentUsecase struct{}

func (m *MockPaymentUsecase) ProcessPayment(payment entities.Payment) error {
	return nil // Simulate success
}

func TestCreatePayment(t *testing.T) {
	usecase := &MockPaymentUsecase{}
	controller := PaymentController{Usecase: usecase}

	router := gin.Default()
	router.POST("/payment", controller.CreatePayment)

	// Create a request body
	body := map[string]interface{}{
		"UserID":           1,
		"PaymentMethodsId": 2,
		"Coin":             100.0,
	}
	bodyBytes, _ := json.Marshal(body)

	// Create a request
	req, _ := http.NewRequest("POST", "/payment", bytes.NewBuffer(bodyBytes))
	w := httptest.NewRecorder()

	// Process the request
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", w.Code)
	}
}
