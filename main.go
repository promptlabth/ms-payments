package main

import (
	"fmt"
	"log"
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/database"
	. "promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/routes"
	"promptlabth/ms-payments/usecases"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

func main() {
	database.DB, err = gorm.Open("postgres", database.DbURL(database.BuildDBConfig()))
	defer database.DB.Close()
	if err != nil {
		log.Fatal("database connect error: ", err)
	} else {
		fmt.Println("connect database successful")

	}
	// auto migrate
	database.DB.AutoMigrate(
		&Coin{},
		&Feature{},
		&Payment{},
		&PaymentMethod{},
		&Feature{},
		&User{},
		&PaymentSubscription{},
	)
	// database.DB.AutoMigrate()

	repo := &repository.PaymentRepository{
		DB: database.DB.DB(),
	}
	usecase := usecases.NewPaymentUsecase(repo)
	controller := controllers.PaymentController{Usecase: usecase}

	r := gin.Default()

	// the clean arch
	routes.CoinRoute(r, database.DB)
	routes.PaymentSubscriptionRoute(r, database.DB)

	r.POST("/payment", controller.CreatePayment)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.Run(":8080")
}
