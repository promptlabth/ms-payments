package main

import (
	"fmt"
	"log"
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/database"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/routes"
	"promptlabth/ms-payments/usecases"

	"gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var err error

func main() {
	database.DB, err = gorm.Open(postgres.Open(database.DbURL(database.BuildDBConfig())), &gorm.Config{})
	// defer database.DB.Close()
	if err != nil {
		log.Fatal("database connect error: ", err)
	} else {
		fmt.Println("connect database successful")

	}
	// auto migrate
	// database.DB.AutoMigrate(
	// 	&Coin{},
	// 	&Feature{},
	// 	&Payment{},
	// 	&PaymentMethod{},
	// 	&Feature{},
	// 	&User{},
	// 	&PaymentSubscription{},
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

	r := gin.Default()

	// the clean arch
	routes.CoinRoute(r, database.DB)
	routes.PaymentSubscriptionRoute(r, database.DB)

	r.POST("/payment", controller.CreatePayment)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.Run()
}
