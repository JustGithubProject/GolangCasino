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

	services.InitNumbersArray(user.TypeOfGame.Numbers)
	money, err := user.TypeOfGame.NormalSpinRoulette(guess_number, bet)
	user.updateBalance(money, bet)
	return user.Balance
}


func (user *UserPlayer) updateBalance(money float64, bet float64){
	if money > bet{
		user.Balance += money
	}else{
		user.Balance -= money
	}
}

func (user *UserPlayer) UnFairPlay(guess_number int, bet float64, gameName string) float64{
	user.TypeOfGame.GameName = gameName
	services.InitWeights(user.TypeOfGame.WeightsForNumbers, 37)
	services.InitWeights(user.TypeOfGame.WeightsForSectors, len(user.TypeOfGame.Sectors))


	services.ShuffleWeights(user.TypeOfGame.WeightsForNumbers)
	services.ShuffleWeights(user.TypeOfGame.WeightsForSectors)


	services.InitNumbersArray(user.TypeOfGame.Numbers)
	money, err := user.TypeOfGame.UnfairSpinRoulette(guess_number, bet)
	user.updateBalance(money, bet)
	return user.Balance
}


