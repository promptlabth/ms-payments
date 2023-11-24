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
	paymentRepo := repository.NewPaymentScriptionRepository(DB)

	userUsecase := usecases.NewUserUseCase(userRepo)
	planUsecase := usecases.NewPlanUsecase(planRepo)
	paymentSubscriptionUsecase := usecases.NewPaymentSubscriptionUsecase(
		paymentRepo,
	)

	subscriptionController := controllers.NewWebhookController(
		userUsecase,
		planUsecase,
		paymentSubscriptionUsecase,
	)

	// for testing (test get invoice ใบเสร็จ from customer id)
	// r.GET("/testing", func(c *gin.Context) {
	// 	listInvoice, _ := services.ListInvoicesByCustomerID("cus_P15cCCU7eDn6c4")
	// 	c.JSON(200, gin.H{
	// 		"data":  listInvoice,
	// 		"count": len(*listInvoice),
	// 	})
	// })

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
			fmt.Println("Subscription canceled!")
			c.JSON(200, gin.H{
				"Test": event,
			})
		default:
			fmt.Printf("Unhandled event type: %s\n", event.Type)
		}
	})
}
