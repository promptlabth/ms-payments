package controllers

import (
	"promptlabth/ms-payments/entities"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockCoinUseCase struct {
}

func (m *MockCoinUseCase) CreateACoin(input *entities.Coin) (err error) {
	return nil // simulate a successful
}
func (m *MockCoinUseCase) DeleteACoin(input *entities.Coin, id string) (err error) {
	return nil // simulate a successful
}
func (m *MockCoinUseCase) UpdateACoin(input *entities.Coin, id string) (err error) {
	return nil // simulate a successful
}
func (m *MockCoinUseCase) GetAllCoins() (coins []entities.Coin, err error) {
	return nil, nil
}
func (m *MockCoinUseCase) GetACoin(input *entities.Coin, id string) (err error) {
	return nil
}

func TestCreateCoin(t *testing.T) {
	usecase := &MockCoinUseCase{}
	controllers := CoinController{coinUseCase: usecase}

	router := gin.Default()
	router.POST("/coin", controllers.CreateACoin)

	// Create request body
	// body := map[string]interface{}{

	// }
}
