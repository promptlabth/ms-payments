package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/promptlabth/ms-payments/controllers"
	"github.com/promptlabth/ms-payments/middlewares"
	"github.com/promptlabth/ms-payments/repository"
	"github.com/promptlabth/ms-payments/usecases"
	"gorm.io/gorm"
)

func SubscriptionRoute(r *gin.Engine, DB *gorm.DB) {

	// initial a repo, usecase, control
	userRepo := repository.NewUserRepository(DB)
	userUseCases := usecases.NewUserUseCase(userRepo)

	planRepo := repository.NewPlanRepository(DB)
	planUsecase := usecases.NewPlanUsecase(planRepo)

	subscriptionController := controllers.NewSubscriptionReqUrlController(
		userUseCases,
		planUsecase,
	)

	// use a middleware to route subscription
	subScription := r.Group("/subscription")
	protect := subScription.Use(middlewares.AuthorizeFirebase())

	protect.POST("/get-url", subscriptionController.GetSubscriptionUrl)

	protect.POST("/get-one-time-url", subscriptionController.GetOneTimeSubscriptionUrl)

	protect.DELETE("/cancle", subscriptionController.CancelSubscription)
}
