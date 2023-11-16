package services

import (
	"context"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func NewFirebaseApplication(c *gin.Context) (*firebase.App, context.Context) {
	// get path of the programs
	home, err := os.Getwd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Not Found a path of server"})
		return nil, nil
	}

	// start connection to firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(home + "/firebase_credential.json")
	config := &firebase.Config{
		ProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
	}
	app, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can't start a firebase service"})
		return nil, nil
	}

	return app, ctx
}
