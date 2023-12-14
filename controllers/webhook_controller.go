package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
	"github.com/stripe/stripe-go/v76"
)

type WebhookController struct {
	userUsecase interfaces.UserUseCase
	planUsecase interfaces.PlanUsecase
}

func NewWebhookController(
	userUsecase interfaces.UserUseCase,
	planUsecase interfaces.PlanUsecase,
) *WebhookController {
	return &WebhookController{
		userUsecase: userUsecase,
		planUsecase: planUsecase,
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
	if paymentSubscription.Status != "active" {
		c.JSON(200, gin.H{
			"err": "สถานะการซื้อขายยังไม่ active กรุณาจ่ายเงินเพื่อให้เราให้สิทธิ์การเข้าใช้งาน subscription",
		})
		return
	}

	customerID := paymentSubscription.Customer.ID
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

	planId := uint(plan.Id)

	user.PlanID = &planId

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": paymentSubscription.CurrentPeriodStart,
		"end":   paymentSubscription.CurrentPeriodEnd,
		"data":  user,
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

	if paymentSubscription.Status != "active" {
		c.JSON(200, gin.H{
			"err": "สถานะการซื้อขายยังไม่ active กรุณาจ่ายเงินเพื่อให้เราให้สิทธิ์การเข้าใช้งาน subscription",
		})
		return
	}
	customerID := paymentSubscription.Customer.ID
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

	newPlanUser := uint(plan.Id)
	user.PlanID = &newPlanUser
	user.Plan = plan

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": paymentSubscription.CurrentPeriodStart,
		"end":   paymentSubscription.CurrentPeriodEnd,
		"data":  user,
		// "data":  paymentSubscription,
	})

}

func (t *WebhookController) DeleteSubscription(c *gin.Context, jsonData []byte) {
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

	var plan entities.Plan
	if err := t.planUsecase.GetAPlan(&plan, 4); err != nil {
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

	newPlanUser := uint(plan.Id)
	user.PlanID = &newPlanUser
	user.Plan = plan

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": paymentSubscription.CurrentPeriodStart,
		"end":   paymentSubscription.CurrentPeriodEnd,
		"data":  user,
		// "data":  paymentSubscription,
	})
}
