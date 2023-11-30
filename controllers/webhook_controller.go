package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
	"github.com/stripe/stripe-go/v76"
)

type WebhookController struct {
	userUsecase                interfaces.UserUseCase
	planUsecase                interfaces.PlanUsecase
	paymentSubscriptionUsecase interfaces.PaymentSubscriptionUseCase
}

func NewWebhookController(
	userUsecase interfaces.UserUseCase,
	planUsecase interfaces.PlanUsecase,
	paymentSubscriptionUsecase interfaces.PaymentSubscriptionUseCase,
) *WebhookController {
	return &WebhookController{
		userUsecase:                userUsecase,
		planUsecase:                planUsecase,
		paymentSubscriptionUsecase: paymentSubscriptionUsecase,
	}
}

func (t *WebhookController) CustromerSubscriptionUpdate(c *gin.Context, jsonData []byte) {

	event := stripe.Event{}
	if err := json.Unmarshal(jsonData, &event); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	var paymentSubscription stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &paymentSubscription); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	customerID := paymentSubscription.Customer.ID
	subscriptionID := paymentSubscription.ID
	productID := paymentSubscription.Items.Data[0].Plan.Product.ID

	var plan entities.Plan
	if err := t.planUsecase.GetAPlanByProdID(&plan, productID); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	var user entities.User
	if err := t.userUsecase.GetAUserByStripeID(&user, customerID); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	// get a data subscription to save a update data
	var prevSubscriptionPayment entities.PaymentSubscription
	if err := t.paymentSubscriptionUsecase.GetSubscriptionPaymentBySubscriptionID(&prevSubscriptionPayment, subscriptionID); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	userId := uint(user.Id)
	planId := uint(plan.Id)
	subscriptionUpdate := entities.PaymentSubscription{
		Id:                 prevSubscriptionPayment.Id,
		SubscriptionID:     subscriptionID,
		StartDatetime:      time.Unix(paymentSubscription.CurrentPeriodStart, 0),
		EndDatetime:        time.Unix(paymentSubscription.CurrentPeriodEnd, 0),
		Datetime:           time.Unix(paymentSubscription.CurrentPeriodStart, 0),
		UserID:             &userId,
		PlanID:             &planId,
		SubscriptionStatus: string(paymentSubscription.Status),
	}

	if err := t.paymentSubscriptionUsecase.UpdateSubscriptionPayment(&subscriptionUpdate); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": paymentSubscription.CurrentPeriodStart,
		"end":   paymentSubscription.CurrentPeriodEnd,
		"data":  subscriptionUpdate,
		// "data":  paymentSubscription,
	})

}

func (t *WebhookController) CreateCustomerSubscription(c *gin.Context, jsonData []byte) {
	// bind a json data to event data
	event := stripe.Event{}
	if err := json.Unmarshal(jsonData, &event); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	// binding a subscription data to variable
	var paymentSubscription stripe.Subscription
	if err := json.Unmarshal(event.Data.Raw, &paymentSubscription); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	customerID := paymentSubscription.Customer.ID
	subscriptionID := paymentSubscription.ID
	productID := paymentSubscription.Items.Data[0].Plan.Product.ID

	var plan entities.Plan
	if err := t.planUsecase.GetAPlanByProdID(&plan, productID); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	var user entities.User
	if err := t.userUsecase.GetAUserByStripeID(&user, customerID); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	userId := uint(user.Id)
	planId := uint(plan.Id)
	subscriptionCreate := entities.PaymentSubscription{
		SubscriptionID:     subscriptionID,
		StartDatetime:      time.Unix(paymentSubscription.CurrentPeriodStart, 0),
		EndDatetime:        time.Unix(paymentSubscription.CurrentPeriodEnd, 0),
		Datetime:           time.Now(),
		UserID:             &userId,
		PlanID:             &planId,
		SubscriptionStatus: string(paymentSubscription.Status),
	}

	if err := t.paymentSubscriptionUsecase.ProcessSubscriptionPayments(&subscriptionCreate); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": paymentSubscription.CurrentPeriodStart,
		"end":   paymentSubscription.CurrentPeriodEnd,
		"data":  subscriptionCreate,
		// "data":  paymentSubscription,
	})

}
