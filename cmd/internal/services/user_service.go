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
	// counter_number_weight := 10
	// counter_sector_weight := 10

	// for i := 0; i < 37; i++{
	// 	user.TypeOfGame.WeightsForNumbers[i] = counter_number_weight
	// 	counter_number_weight += 100
	// }
	
	// for i := 0; i < len(user.TypeOfGame.Sectors); i++{
	// 	user.TypeOfGame.WeightsForSectors[i] = counter_sector_weight
	// 	counter_sector_weight += 100
	// }

	services.InitWeights(user.TypeOfGame.WeightsForNumbers, 37)
	services.InitWeights(user.TypeOfGame.WeightsForSectors, len(user.TypeOfGame.Sectors))


	services.ShuffleWeights(user.TypeOfGame.WeightsForNumbers)
	services.ShuffleWeights(user.TypeOfGame.WeightsForSectors)


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


