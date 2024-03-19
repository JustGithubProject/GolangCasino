package services

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/services"
)

type UserPlayer struct {
	TypeOfGame services.GameRoulette
	Balance    int
}

func (user *UserPlayer) Play(guess_number int, bet int, gameName string) int {
	user.TypeOfGame.GameName = gameName
	// user.TypeOfGame.Numbers = []int{0, 3, 6, 9, 12, 2, 5, 8, 11, 1, 4, 7, 10, 15, 18, 21, 24, 14, 17, 20, 23, 13, 16, 19, 22}
	currentNumber := 0

	for i := 0; i < 36; i++{
		user.TypeOfGame.Numbers[i] = currentNumber
		currentNumber += 3
	}
	
	money, err := user.TypeOfGame.SpinRoulette(guess_number, bet)
	if money > bet {
		user.Balance += money
	} else {
		user.Balance -= money
	}
	return user.Balance
}
