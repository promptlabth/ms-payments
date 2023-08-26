package entities

import (
	"time"
)

type Feature struct {
	Id           int `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	DateOfCreate time.Time
	Url          string

	Payment []Payment `gorm:"foreignKey:FeatureID"`

}
