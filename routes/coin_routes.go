package routes

import (
	"promptlabth/ms-payments/controllers"
	"promptlabth/ms-payments/repository"
	"promptlabth/ms-payments/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CoinRoute(r *gin.Engine, DB *gorm.DB) {
	coinRepo := repository.NewCoinRepository(DB)
	coinUseCase := usecases.NewCoinUseCase(coinRepo)
	coinController := controllers.NewCoinController(coinUseCase)

	r.GET("coin", coinController.GetAllCoins)
	r.GET("coin/:id", coinController.GetACoin)
	r.POST("coin", coinController.CreateACoin)
	r.PATCH("coin/:id", coinController.UpdateACoin)
}
