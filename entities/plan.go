// entities/plan.go
package entities

type Plan struct {
	Id       int `gorm:"primaryKey;autoIncrement:true"`
	PlanType string
	Datetime string

	PaymentSubscriptions []PaymentSubscription `gorm:"foreignKey:PlanID"`
}

func (b *Plan) TableName() string {
	return "plans"
}
