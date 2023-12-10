package main

import (
	"log"
	"os"

	"github.com/promptlabth/ms-payments/database"
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/promptlabth/ms-payments/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}

var err error

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /api
func main() {
	database.DB, err = gorm.Open(postgres.Open(
		database.DbURL(database.BuildDBConfig()),
	), &gorm.Config{})
	// defer database.DB.Close()
	if err != nil {
		log.Fatal("database connect error: ", err)
	}
	// auto migrate
	database.DB.AutoMigrate(
		&entities.Coin{},
		&entities.User{},
		&entities.Plan{},
	)
	// database.DB.AutoMigrate()
	if os.Getenv("BaseOn") != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// to set CORS
	r.Use(CORSMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	if err != nil {
		log.Fatal(err)
	}

	// the clean arch

	routes.CoinRoute(r, database.DB)

	routes.SubscriptionRoute(r, database.DB)

	routes.WebhookRoute(r, database.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r.Run(":" + port)
}
