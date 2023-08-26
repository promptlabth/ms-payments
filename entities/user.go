package entities

type User struct{
	Id int `gorm:"primaryKey;autoIncrement:true"`
	Firebase_id string
	Name string
	Email string
	ProfilePic string

	Payment []Payment `gorm:"foreignKey:UserID"`
	Coin []Coin `gorm:"foreignKey:UserID"`
}