package services


import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/services"
)

type UserPlayer struct {
	TypeOfGame services.GameRoulette
	Balance int
}


func (user *UserPlayer) Play(guess_number int, bet int) int{
	money, err := user.TypeOfGame.SpinRoulette(guess_number, bet)
	if money > bet{
		user.Balance += money
	}else{
		user.Balance -= money
	}
	return user.Balance
}