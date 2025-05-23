package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"time"

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
	startTime := time.Unix(paymentSubscription.CurrentPeriodStart, 0)
	endTime := time.Unix(paymentSubscription.CurrentPeriodStart, 0).AddDate(0, 1, 0)

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
	user.Sub_date = startTime
	user.End_sub_date = endTime
	user.Monthly = true

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": startTime.Unix(),
		"end":   endTime.Unix(),
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
	startTime := time.Unix(paymentSubscription.CurrentPeriodStart, 0)
	endTime := time.Unix(paymentSubscription.CurrentPeriodStart, 0).AddDate(0, 1, 0)

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
	user.Sub_date = startTime
	user.End_sub_date = endTime
	user.Monthly = true

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": startTime.Unix(),
		"end":   endTime.Unix(),
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
	user.Sub_date = time.Time{}
	user.End_sub_date = time.Time{}
	user.Monthly = false

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

func (t *WebhookController) OneTimeCustomerSubscription(c *gin.Context, jsonData []byte) {
	event := stripe.Event{}
	if err := json.Unmarshal(jsonData, &event); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	var paymentSubscription stripe.PaymentIntent
	if err := json.Unmarshal(event.Data.Raw, &paymentSubscription); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}
	if paymentSubscription.Status != "succeeded" {
		c.JSON(200, gin.H{
			"err": "สถานะการซื้อขายยังไม่สำเร็จ กรุณาจ่ายเงินเพื่อให้เราให้สิทธิ์การเข้าใช้งาน subscription",
		})
		return
	}

	customerID := paymentSubscription.Customer.ID
	planPrice := int(paymentSubscription.Amount) / 100 // convert to baht (from 5900 to decimal 59.00)
	startTime := time.Unix(paymentSubscription.Created, 0)
	endTime := time.Unix(paymentSubscription.Created, 0).AddDate(0, 1, 0)

	var plan entities.Plan
	if err := t.planUsecase.GetAPlanByPrice(&plan, planPrice); err != nil {
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
	user.Sub_date = startTime
	user.End_sub_date = endTime
	user.Monthly = false

	if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"start": startTime.Unix(),
		"end":   endTime.Unix(),
		"data":  user,
		// "data":  paymentSubscription,
	})
}
