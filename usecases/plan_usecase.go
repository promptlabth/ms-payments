package usecases

import (
	"github.com/promptlabth/ms-payments/entities"
	"github.com/promptlabth/ms-payments/interfaces"
)

type planUsecase struct {
	planRepo interfaces.PlanRepository
}

func NewPlanUsecase(planRepo interfaces.PlanRepository) interfaces.PlanUsecase {
	return &planUsecase{
		planRepo: planRepo,
	}
}

func (t *planUsecase) GetAPlan(plan *entities.Plan, id int) error {
	if err := t.planRepo.GetAPlan(plan, id); err != nil {
		return err
	}
	return nil
}

func (t *planUsecase) GetAPlanByPriceID(plan *entities.Plan, id string) error {
	if err := t.planRepo.GetAPlanByPriceID(plan, id); err != nil {
		return nil
	}
	return nil
}

func (t *planUsecase) GetAPlanByProdID(plan *entities.Plan, id string) error {
	if err := t.planRepo.GetAPlanByProdID(plan, id); err != nil {
		return err
	}
	return nil
}

func (t *planUsecase) CreateAPlan(plan *entities.Plan) error {
	if err := t.planRepo.CreateAPlan(plan); err != nil {
		return err
	}
	return nil
}

func (t *planUsecase) GetAPlanByPrice(plan *entities.Plan, price int) error {
	if err := t.planRepo.GetAPlanByPrice(plan, price); err != nil {
		return err
	}
	return nil
}
