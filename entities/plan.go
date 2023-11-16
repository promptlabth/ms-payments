// entities/plan.go
package entities

type Plan struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	PlanType    string `gorm:"column:planType"`
	Datetime    string
	MaxMessages int `gorm:"column:maxMessages"`

	PaymentSubscriptions []PaymentSubscription `gorm:"foreignKey:PlanID"`
}

func (b *Plan) TableName() string {
	return "plans"
}
