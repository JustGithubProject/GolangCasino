package models 

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name string
	Email string `gorm."unique"`
	Password string
	Balance float64
	Payments []Payment
}


type Payment struct {
	gorm.Model
	OrderID string
	UserID uint
	Status string
}