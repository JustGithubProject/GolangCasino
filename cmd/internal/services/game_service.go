package services


import (
    "fmt"
    "math/rand"
)

type GameRoulette struct{
	GameName string
	Numbers []int
	Sectors []string
}

func (game *GameRoulette) ChooseRandomNumber(array []int) int {
    // Generate a random index within the range of the array length
    randomIndex := rand.Intn(len(array))
    return array[randomIndex]
}

func (game *GameRoulette) ChooseRandomSector(sectors []string) string{
	randomIndex := rand.Intn(len(sectors))
	return sectors[randomIndex]
}

func (game *GameRoulette) GenerateRandomNumberFromArray(numbers []int) int {
    return game.ChooseRandomNumber(numbers)
}

func (game *GameRoulette) GenerateRandomSectorFromArray(sectors []string) string{
	return game.ChooseRandomSector(sectors)
}

func (game *GameRoulette) SpinRoulette(sector string, guess_number int, bet int) (int, error){
	if sector != nil{
		dropped_sector := game.GenerateRandomSectorFromArray(game.Sectors)
		if dropped_sector == sector{
			prize := bet * 3
			fmt.Printf("You won %d\n", prize)
			return prize, nil
		}
	}
	dropped_number := game.GenerateRandomNumberFromArray(game.Numbers)
	if dropped_number == guess_number{
		prize := bet * 2
		fmt.Printf("You won %d\n", prize)
		return prize, nil
	}else{
		fmt.Printf("You lose %d\n", bet)
		return bet, nil
	}
}

