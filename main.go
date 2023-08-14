package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/usecases"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := &repository.PaymentRepository{
		DB: db,
	}
	usecase := usecases.NewPaymentUsecase(repo)
	controller := controllers.PaymentController{Usecase: usecase}

	r := gin.Default()
	r.POST("/payment", controller.CreatePayment)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	r.Run(":8080")
}
