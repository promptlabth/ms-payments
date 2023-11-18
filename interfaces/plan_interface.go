package interfaces

import "github.com/promptlabth/ms-payments/entities"

type PlanRepository interface {
	GetAPlan(plan *entities.Plan, id int) error
	GetAPlanByPriceID(plan *entities.Plan, id string) error
	CreateAPlan(plan *entities.Plan) error
}

type PlanUsecase interface {
	GetAPlan(plan *entities.Plan, id int) error
	GetAPlanByPriceID(plan *entities.Plan, id string) error
	CreateAPlan(plan *entities.Plan) error
}
