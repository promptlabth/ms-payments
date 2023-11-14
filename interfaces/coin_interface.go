package interfaces

import "github.com/promptlabth/ms-payments/entities"

type CoinUseCase interface {
	GetAllCoins() (t []entities.Coin, err error)
	CreateACoin(t *entities.Coin) (err error)
	GetACoin(t *entities.Coin, id string) (err error)
	UpdateACoin(t *entities.Coin, id string) (err error)
	DeleteACoin(t *entities.Coin, id string) (err error)
}

type CoinRepository interface {
	GetAllCoins(t *[]entities.Coin) (err error)
	CreateACoin(t *entities.Coin) (err error)
	GetACoin(t *entities.Coin, id string) (err error)
	UpdateACoin(t *entities.Coin, id string) (err error)
	DeleteACoin(t *entities.Coin, id string) (err error)
}
