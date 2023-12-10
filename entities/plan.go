// entities/plan.go
package entities

type Plan struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	PlanType    string `gorm:"column:planType"`
	MaxMessages int    `gorm:"column:maxMessages"`
	ProductID   string `gorm:"uniqueIndex:idx_plan_product_id"`

	Users []User `gorm:"foreignKey:PlanID"`
}

func (b *Plan) TableName() string {
	return "plans"
}
