// entities/plan.go
package entities

type Plan struct {
	Id                  int `gorm:"primaryKey;autoIncrement:true"`
	PlanType 			string
	Datetime            string
}

func (b *Plan) TableName() string {
	return "plan"
}
