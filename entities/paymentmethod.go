package entities

type PaymentMethod struct{
	Id int `gorm:"primaryKey;autoIncrement:true"`
	Type string

	Payments []Payment `gorm:"foreignKey:PaymentMethodID"`
}