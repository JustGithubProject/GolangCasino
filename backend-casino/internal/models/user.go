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

	GoogleID string `gorm."unique"`
	Picture string  
	GivenName string   // Name
	FamilyName string  // Surname
	Locale string 
}


type Payment struct {
	gorm.Model
	OrderID string
	UserID uint
	Amount float64
	Status string
}