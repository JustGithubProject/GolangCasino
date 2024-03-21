package services

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/services"
)

type UserPlayer struct {
	TypeOfGame services.GameRoulette
	Balance		float64
}

func (user *UserPlayer) Play(guess_number int, bet int, gameName string) float64 {
	user.TypeOfGame.GameName = gameName

	for i := 0; i < 37; i++{
		user.TypeOfGame.Numbers[i] = i
	}

	money, err := user.TypeOfGame.SpinRoulette(guess_number, bet)
	if money > bet {
		user.Balance += money
	} else {
		user.Balance -= money
	}
	return user.Balance
}



