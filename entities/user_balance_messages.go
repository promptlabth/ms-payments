package entities

type BalanceMesssage struct {
	Message int `gorm:"column:balance_message"`

	FirebaseID string `gorm:"column:firebase_id"`
	User       User   `gorm:"foreignKey:FirebaseID;references:Firebase_id"`
}

func (b *BalanceMesssage) TableName() string {
	return "user_balance_messages"
}
