package usecases

import (
	"promptlabth/ms-payments/entities"
	"promptlabth/ms-payments/interfaces"
)

type coinUseCase struct {
	coinRepo interfaces.CoinRepository
}

func NewCoinUseCase(repo interfaces.CoinRepository) interfaces.CoinUseCase {
	return &coinUseCase{
		coinRepo: repo,
	}
}

func (c *coinUseCase) GetAllCoins() (coins []entities.Coin, err error) {
	var coin []entities.Coin
	handleErr := c.coinRepo.GetAllCoins(&coin)
	return coin, handleErr
}

func (c *coinUseCase) CreateACoin(input *entities.Coin) (err error) {
	handleErr := c.coinRepo.CreateACoin(input)
	return handleErr
}

func (c *coinUseCase) GetACoin(input *entities.Coin, id string) (err error) {
	handleErr := c.coinRepo.GetACoin(input, id)
	return handleErr
}

func (c *coinUseCase) UpdateACoin(input *entities.Coin, id string) (err error) {
	// check is avaliable
	var checkingCoin entities.Coin
	errRes := c.coinRepo.GetACoin(&checkingCoin, id)
	if errRes != nil {
		return errRes
	}
	// update
	handleErr := c.coinRepo.UpdateACoin(input, id)
	return handleErr
}

func (c *coinUseCase) DeleteACoin(input *entities.Coin, id string) (err error) {
	handleErr := c.coinRepo.DeleteACoin(input, id)
	return handleErr
}
