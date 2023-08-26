package entities

type Coin struct {
	Id    int `gorm:"primaryKey;autoIncrement:true"`
	Total int32

	// user
	UserID *int `valid:"-"`
	User   User `gorm:"references:Id" valid:"-"`
}
