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
}

func NewSubscriptionReqUrlController(
	userusecase interfaces.UserUseCase,
) *SubscriptionReqUrlController {
	return &SubscriptionReqUrlController{
		userUsecase: userusecase,
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
