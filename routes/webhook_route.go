package routes

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/controllers"
	"github.com/promptlabth/ms-payments/repository"
	"github.com/promptlabth/ms-payments/usecases"
	"github.com/stripe/stripe-go/v76"
	"gorm.io/gorm"
)

type StripeData struct {
	status string
}

func WebhookRoute(r *gin.Engine, DB *gorm.DB) {
	userRepo := repository.NewUserRepository(DB)
	planRepo := repository.NewPlanRepository(DB)

	userUsecase := usecases.NewUserUseCase(userRepo)
	planUsecase := usecases.NewPlanUsecase(planRepo)

	subscriptionController := controllers.NewWebhookController(
		userUsecase,
		planUsecase,
	)

	r.POST("/webhook", func(c *gin.Context) {

		jsonData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return
		}

		event := stripe.Event{}
		if err := json.Unmarshal(jsonData, &event); err != nil {
			return
		}

		switch event.Type {
		case "invoice.payment_succeeded":
			// Handle success payment
			fmt.Println("Payment succeeded!")
		case "invoice.payment_failed":
			// handle payment failure
			fmt.Println("Payment failed!")
		case "customer.subscription.created":
			// Handle subscription creation
			fmt.Println("================ EVENT customer.subscription.created ======================", event.Object)
			subscriptionController.CreateCustomerSubscription(c, jsonData)

		case "customer.subscription.updated":
			// Handle subscription renewal or active start subscription
			fmt.Println("================ EVENT customer.subscription.Update ======================", event.Object)
			subscriptionController.CustromerSubscriptionUpdate(c, jsonData)

		case "customer.subscription.deleted":
			// Handle subscription cancellation
			subscriptionController.DeleteSubscription(c, jsonData)
			c.JSON(200, gin.H{
				"Test": event,
			})
		default:
			fmt.Printf("Unhandled event type: %s\n", event.Type)
		}
	})
}
