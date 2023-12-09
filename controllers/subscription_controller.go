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
	userUsecase interfaces.UserUseCase
	planUsecase interfaces.PlanUsecase
}

func NewSubscriptionReqUrlController(
	userusecase interfaces.UserUseCase,
	planUsecase interfaces.PlanUsecase,
) *SubscriptionReqUrlController {
	return &SubscriptionReqUrlController{
		userUsecase: userusecase,
		planUsecase: planUsecase,
	}
}

// Request Struct Input
type SubscriptionReqUrl struct {
	PrizeID string
	WebUrl  string
	PlanID  int
}

func (t *SubscriptionReqUrlController) GetSubscriptionUrl(c *gin.Context) {

	// checkout session to stripe service
	var subscriptionReqUrl SubscriptionReqUrl
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

	// check a customer is have a subscription? (when user click a same plan [other plan mean upgrade])
	if *user.PlanID != uint(4) {
		c.JSON(404, gin.H{
			"error": "คุณไม่สามารถ subscription อีกครั้งได้",
		})
		return
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

// struct for cancel subscription (only subscription id)
type CancelSubscriptionRequest struct {
	SubscriptionID string
}

func (t *SubscriptionReqUrlController) CancelSubscription(c *gin.Context) {

	firebaseUID := c.GetString("firebase_id")

	var user entities.User
	// get a user by firebase id
	if err := t.userUsecase.GetAUserByFirebaseId(&user, firebaseUID); err != nil {
		// if not found a user from firebase id
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	subscriptionsAddress, err := services.ListSubscriptionByCustomerID(user.StripeId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบข้อมูล subscription",
		})
		return
	}
	subscriptions := *subscriptionsAddress

	// cancel subscription At period end
	subscription, err := services.CancelAtPeriodBySubID(subscriptions[0].ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200,
		gin.H{
			"data": subscription,
		})

}
