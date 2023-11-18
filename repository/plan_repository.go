package repository

import (
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
	"gorm.io/gorm"
)

type planRepository struct {
	conn *gorm.DB
}

func NewPlanRepository(conn *gorm.DB) interfaces.PlanRepository {
	return &planRepository{
		conn: conn,
	}
}

func (t *planRepository) GetAPlan(plan *entities.Plan, id int) error {
	if err := t.conn.Where("id = ?", id).First(plan).Error; err != nil {
		return err
	}
	return nil
}

func (t *planRepository) GetAPlanByPriceID(plan *entities.Plan, id string) error {
	if err := t.conn.Where("price_id = ?", id).First(plan).Error; err != nil {
		return err
	}
	return nil
}

func (t *planRepository) CreateAPlan(plan *entities.Plan) error {
	if err := t.conn.Create(plan).Error; err != nil {
		return err
	}
	return nil
}
