package main

import (
	"fmt"
	"log"
	"os"
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/database"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/routes"
	"promptlabth/ms-payments/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func main() {
	database.DB, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_USER"), os.Getenv("DB_NAME")),
	), &gorm.Config{})
	// defer database.DB.Close()
	if err != nil {
		log.Fatal("database connect error: ", err)
	}
	// auto migrate
	// database.DB.AutoMigrate(
	// 	&entities.Coin{},
	// 	&entities.Feature{},
	// 	&entities.Payment{},
	// 	&entities.PaymentMethod{},
	// 	&entities.Feature{},
	// 	&entities.User{},
	// 	&entities.PaymentSubscription{},
	// )
	// database.DB.AutoMigrate()

	repo := &repository.PaymentRepository{}
	db, err := database.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	repo.DB = db
	usecase := usecases.NewPaymentUsecase(repo)
	controller := controllers.PaymentController{Usecase: usecase}
	if os.Getenv("BaseOn") != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// to set CORS
	r.Use(CORSMiddleware())
	// the clean arch
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	routes.CoinRoute(r, database.DB)
	routes.PaymentSubscriptionRoute(r, database.DB)

	r.POST("/payment", controller.CreatePayment)

	// Health check endpoint

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r.Run(":" + port)
}
