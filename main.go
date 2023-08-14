package main

import (
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/usecases"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.PaymentRepository{DB: db}
	usecase := usecases.PaymentUsecase{Repository: &repo}
	controller := controllers.PaymentController{Usecase: usecase}

	r := gin.Default()
	r.POST("/payment", controller.CreatePayment)
	r.Run(":8080")
}
