// entities/plan.go
package entities

type Plan struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	PlanType    string `gorm:"column:planType"`
	MaxMessages int    `gorm:"column:maxMessages"`
	PriceID     string `gorm:"uniqueIndex:idx_plan_price_id"`

	PaymentSubscriptions []PaymentSubscription `gorm:"foreignKey:PlanID"`
}

func (b *Plan) TableName() string {
	return "plans"
}
