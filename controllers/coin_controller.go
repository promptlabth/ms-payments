package controllers

import (
	"net/http"
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"

	"github.com/gin-gonic/gin"
)

type CoinController struct {
	coinUseCase interfaces.CoinUseCase
}

func NewCoinController(usecase interfaces.CoinUseCase) *CoinController {
	return &CoinController{
		coinUseCase: usecase,
	}
}

func (t *CoinController) GetAllCoins(c *gin.Context) {
	resp, err := t.coinUseCase.GetAllCoins()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (t *CoinController) GetACoin(c *gin.Context) {
	var coin entities.Coin
	if err := t.coinUseCase.GetACoin(&coin, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, coin)
	}
}

func (t *CoinController) CreateACoin(c *gin.Context) {
	var newCoin entities.Coin
	if err := c.ShouldBindJSON(&newCoin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := t.coinUseCase.CreateACoin(&newCoin); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, newCoin)
	}
}

func (t *CoinController) UpdateACoin(c *gin.Context) {
	var newCoin entities.Coin
	if err := c.ShouldBindJSON(&newCoin); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := t.coinUseCase.UpdateACoin(&newCoin, c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
}