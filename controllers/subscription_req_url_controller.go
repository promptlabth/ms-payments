package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
	"github.com/promptlabth/ms-payments/services"
)

type SubscriptionReqUrlController struct {
	userUsecase                interfaces.UserUseCase
	planUsecase                interfaces.PlanUsecase
	paymentSubscriptionUsecase interfaces.PaymentSubscriptionUseCase
}

func NewSubscriptionReqUrlController(
	userusecase interfaces.UserUseCase,
	planUsecase interfaces.PlanUsecase,
	paymentSubscriptionUsecase interfaces.PaymentSubscriptionUseCase,
) *SubscriptionReqUrlController {
	return &SubscriptionReqUrlController{
		userUsecase:                userusecase,
		planUsecase:                planUsecase,
		paymentSubscriptionUsecase: paymentSubscriptionUsecase,
	}
}
func (t *SubscriptionReqUrlController) GetSubscriptionUrl(c *gin.Context) {

	// checkout session to stripe service
	var subscriptionReqUrl entities.SubscriptionReqUrl
	if err := c.ShouldBindJSON(&subscriptionReqUrl); err != nil {
		// Respond with a 400 Bad Request status code and detailed error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get a firebase UID from gin context
	firebaseUID := c.GetString("firebase_id")

	var user entities.User

	// use firebase UID to find is found in database
	if err := t.userUsecase.GetAUserByFirebaseId(&user, firebaseUID); err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check is have a customer stripe id?
	if user.StripeId == "" {
		// if not found a stripe id will be create a customer stripe id to this user and update it

		// create a customer in stripe
		cus, err := services.CreateCustomer(user.Email, user.Name, user.Firebase_id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// update a user in stripe
		user.StripeId = cus.ID
		if err := t.userUsecase.UpdateAUser(&user, strconv.Itoa(user.Id)); err != nil {
			if err != nil {
				c.JSON(403, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}

	// to create a cehckout url from stripe (make subscription url to customer)
	checkoutSession, err := services.CreateCheckoutSession(
		subscriptionReqUrl.PrizeID,
		"subscription",
		[]string{
			"card",
		},
		user.StripeId,
		subscriptionReqUrl.WebUrl,
		subscriptionReqUrl.PlanID,
	)
	if err != nil {
		// if error found will return can't checkout session in backend
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if everything is successful will send a url to checkout to frontend
	c.JSON(201, gin.H{
		"url": checkoutSession.URL,
	})

}

// for save a data subscription to payment subscription
func (t *SubscriptionReqUrlController) SaveSubscription(c *gin.Context) {
	// get a firebase id from middleware
	firebaseId := c.GetString("firebase_id")

	// binding a requset from user
	var subscriptionReq entities.SaveSubscriptionReq
	if err := c.ShouldBindJSON(&subscriptionReq); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// find a checkout data from checkout session id
	session, err := services.CheckCheckoutSessionId(subscriptionReq.CheckoutSessionId)

	// if found a error in check session id
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check a status is complete?
	if session.Status != "complete" {
		c.JSON(400, gin.H{
			"error": "คุณยังจ่ายเงินไม่สำเร็จ",
		})
		return
	}

	// check a session is success?
	if session.PaymentStatus != "paid" {
		c.JSON(400, gin.H{
			"error": "คุณยังไม่ได้จ่ายเงินให้กับเรา",
		})
		return
	}

	// find a payment subscription haved?
	var oldPayment entities.PaymentSubscription
	if err := t.paymentSubscriptionUsecase.GetSubscriptionPaymentBySubscriptionID(&oldPayment, session.Subscription.ID); err == nil {
		c.JSON(400, gin.H{
			"error": "คุณเคยใช้ ID นี้จ่ายเงินแล้ว หากไม่ได้สิทธิ์กรุณาติดต่อแอดมิน",
		})
		return
	}

	// find a user by firebase id
	var user entities.User
	if err := t.userUsecase.GetAUserByFirebaseId(&user, firebaseId); err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบข้อมูล User รายนี้",
		})
		return
	}

	// find a price by subscription id from stripe
	subscriptionItem, _ := services.GetPriceBySubscriptionID(session.Subscription.ID)

	// find a plan by price id
	var plan entities.Plan
	if err := t.planUsecase.GetAPlanByPriceID(&plan, subscriptionItem.Plan.ID); err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบข้อมูล plan นี้",
		})
		return
	}

	// add a subscription id to payment subscriptions
	userId := uint(user.Id)
	planId := uint(plan.Id)
	ps := entities.PaymentSubscription{
		SubscriptionID:  session.Subscription.ID,
		UserID:          &userId,
		PlanID:          &planId,
		PaymentMethodID: nil,
	}
	if err := t.paymentSubscriptionUsecase.ProcessSubscriptionPayments(ps); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"data":             session,
		"user":             user,
		"plan":             plan,
		"session":          session,
		"subscriptionItem": subscriptionItem.Plan,
	})

}
