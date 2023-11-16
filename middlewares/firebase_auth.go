package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/services"
)

func AuthorizeFirebase() gin.HandlerFunc {
	return func(c *gin.Context) {

		// initial firebase application
		app, firebaseCtx := services.NewFirebaseApplication(c)

		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Format of Authorization Token"})
			return
		}

		client, err := app.Auth(firebaseCtx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// extract a client token to token data
		token, err := client.VerifyIDToken(firebaseCtx, clientToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// set a data to gin context to save a firebase_ID (FirebaseUID)
		c.Set("firebase_id", token.UID)
		c.Next()

	}
}
