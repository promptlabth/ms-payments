package routes

import (
	"github.com/promptlabth/ms-payments/controllers"
	"github.com/promptlabth/ms-payments/repository"
	"github.com/promptlabth/ms-payments/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PaymentSubscriptionRoute(r *gin.Engine, DB *gorm.DB) {
	// 1. Initialize the PaymentSubscriptionsRepository
	repo := repository.NewPaymentScriptionRepository(DB)

	// 2. Initialize the PaymentSubscriptionUsecase
	paymentSubscriptionUsecase := usecases.NewPaymentSubscriptionUsecase(repo)

	// 3. Initialize the PaymentSubscriptionController
	paymentSubscriptionController := controllers.NewPaymentScriptionController(paymentSubscriptionUsecase)

	// 4. Define the routes and associate them with the controller methods
	r.POST("/payment-subscription", paymentSubscriptionController.CreatePaymentSubscription)
	// Add more routes as needed
}
