package repository

import (
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type coinRepository struct {
	conn *gorm.DB
}

func NewCoinRepository(connection *gorm.DB) interfaces.CoinRepository {
	return &coinRepository{
		conn: connection,
	}
}

func (t *coinRepository) GetAllCoins(coin *[]entities.Coin) (err error) {
	if err := t.conn.Find(coin).Error; err != nil {
		return err
	}
	return nil
}

func (t *coinRepository) CreateACoin(coin *entities.Coin) (err error) {
	if err := t.conn.Create(coin).Error; err != nil {
		return err
	}
	return nil
}

func (t *coinRepository) GetACoin(coin *entities.Coin, id string) (err error) {
	if err := t.conn.Where("id = ?", id).First(coin).Error; err != nil {
		return err
	}
	return nil
}

func (t *coinRepository) UpdateACoin(coin *entities.Coin, id string) (err error) {
	t.conn.Save(coin)
	return nil
}

func (t *coinRepository) DeleteACoin(coin *entities.Coin, id string) (err error) {
	t.conn.Where("id = ?", id).Delete(coin)
	return nil
}
