package entities

type User struct {
	Id          int    `gorm:"primaryKey;autoIncrement:true"`
	Firebase_id string `gorm:"uniqueIndex:idx_user_firebase_id"`
	Name        string
	Email       string
	Profilepic  string
	StripeId    string `gorm:"uniqueIndex:idx_user_stripe_id"`

	Payment []Payment `gorm:"foreignKey:UserID"`
	Coin    []Coin    `gorm:"foreignKey:UserID"`

	PaymentSubscriptions []PaymentSubscription `gorm:"foreignKey:UserID"`
}

func (b *User) TableName() string {
	return "users"
}
