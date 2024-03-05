package models

import (
	"time"

	"gorm.io/gorm"
)


type Game struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignKey:UserID"`
	GameType string
	BetAmount float64
	WinAmount float64
	PlayedAt time.Time
}