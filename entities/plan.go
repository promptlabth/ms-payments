// entities/plan.go
package entities

type Plan struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	PlanType    string `gorm:"column:plan_type"`
	MaxMessages int    `gorm:"column:max_messages"`
	ProductID   string `gorm:"uniqueIndex:idx_plan_product_id"`

	Users []User `gorm:"foreignKey:PlanID"`
}

func (b *Plan) TableName() string {
	return "plans"
}
