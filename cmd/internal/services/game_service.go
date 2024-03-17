package services


import (
    "fmt"
    "math/rand"
    "time"
)

type GameRoulette struct{
	GameName string
	Numbers []int
}

func (game *GameRoulette) ChooseRandomNumber(array []int) int {
    // Generate a random index within the range of the array length
    randomIndex := rand.Intn(len(array))
    return array[randomIndex]
}

func (game *GameRoulette) GenerateRandomFromArray(numbers []int) int {
    return game.ChooseRandomNumber(numbers)
}

func (game *GameRoulette) SpinRoulette(guess_number int, bet int) (int, error){
	dropped_number := game.GenerateRandomFromArray(game.Numbers)
	if dropped_number == guess_number{
		prize := bet * 2
		fmt.Printf("You won %d\n", prize)
		return prize, nil
	}else{
		fmt.Printf("You lose %d\n", bet)
		return bet, nil
	}
}

