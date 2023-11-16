package entities

type User struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	Firebase_id string `gorm:"uniqueKey"`
	Name        string
	Email       string
	Profilepic  string
	StripeId    string `gorm:"uniqueKey"`

	Payment []Payment `gorm:"foreignKey:UserID"`
	Coin    []Coin    `gorm:"foreignKey:UserID"`

	PaymentSubscriptions []PaymentSubscription `gorm:"foreignKey:UserID"`
}

func (b *User) TableName() string {
	return "users"
}
