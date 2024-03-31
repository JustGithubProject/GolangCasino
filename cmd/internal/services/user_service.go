package services

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/services"
	"math/rand"
	"time"
)

type UserPlayer struct {
	TypeOfGame services.GameRoulette
	Balance		float64
}

func (user *UserPlayer) NormalPlay(guess_number int, bet int, gameName string) float64 {
	user.TypeOfGame.GameName = gameName

	for i := 0; i < 37; i++{
		user.TypeOfGame.Numbers[i] = i
	}

	money, err := user.TypeOfGame.NormalSpinRoulette(guess_number, bet)
	if money > bet {
		user.Balance += money
	} else {
		user.Balance -= money
	}
	return user.Balance
}


func (user *UserPlayer) UnFairPlay(guess_number int, bet int, gameName string) float64{
	user.TypeOfGame.GameName = gameName
	counter_number_weight := 10
	counter_sector_weight := 10

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 37; i++{
		user.TypeOfGame.WeightsForNumbers[i] = counter_number_weight
		counter_number_weight += 100
	}
	
	for i := 0; i < len(user.TypeOfGame.Sectors); i++{
		user.TypeOfGame.WeightsForSectors[i] = counter_sector_weight
		counter_sector_weight += 100
	}

	services.shuffle_weights(user.TypeOfGame.WeightsForNumbers)
	services.shuffle_weights(user.TypeOfGame.WeightsForSectors)


	for i := 0; i < 37; i++{
		user.TypeOfGame.Numbers[i] = i
	}
	money, err := user.TypeOfGame.UnfairSpinRoulette(guess_number, bet)
	if money > bet {
		user.Balance += money
	} else {
		user.Balance -= money
	}
	return user.Balance
}


