package services


import (
    "math/rand"
)


type GameRoulette struct{
	GameName string
	Numbers []int
	Sectors []string

    WeightsForNumbers []int
}

//////////////////////////////////////////////////////////////////////////////////////////////
//              UNFAIRSPINROULETTE   (LOOK DOWN)                                			//
//////////////////////////////////////////////////////////////////////////////////////////////

func (game *GameRoulette) CheckColor(number int) string{
	redArray := [18]int{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
	blackArray := [18]int{2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35}
	
	// Is red?
	for i := 0; i < len(redArray); i++{
		if number == redArray[i]{
			return "red"
		}
	}
	// Is black?
	for i := 0; i < len(blackArray); i++{
		if number == blackArray[i]{
			return "black"
		}
	}
	// Else dropped 0 that means green
	return "green"
}


func (game *GameRoulette) ChooseRandomNumberByWeight(numbers []int, weights []int) int {
    totalWeight := 0
    for _, weight := range weights {
        totalWeight += weight
    }

    r := rand.Intn(totalWeight)

    cumulativeWeight := 0
    for i, weight := range weights {
        cumulativeWeight += weight
        if r < cumulativeWeight {
            return numbers[i]
        }
    }

    // This should never happen if weights are correctly provided
    return numbers[len(numbers)-1]
}



func (game *GameRoulette) GenerateRandomNumberByWeight(numbers []int, weights []int) int{
    return game.ChooseRandomNumberByWeight(numbers, weights)
}


func (game *GameRoulette) UnfairSpinRoulette(redToBets map[string]float64, blackToBets map[string]float64, sectorsToBets map[string]float64, numbersToBets map[int]float64) (float64, error){
	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	lengthOfBetsToRed := len(redToBets)
	lengthOfBetsToBlack := len(blackToBets)


	dropped_number := game.GenerateRandomNumberByWeight(game.Numbers, game.WeightsForNumbers)
	prize := float64(0)

	if lengthOfBetsToNumbers > 0{
		if _, ok := numbersToBets[dropped_number]; ok{
			prize += numbersToBets[dropped_number] * float64(35)
		}
	}
	if lengthOfBetsToSectors > 0{
		dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
		if _, ok := sectorsToBets[dropped_sector]; ok{
			prize += sectorsToBets[dropped_sector] * float64(3)
		}
	}
	color := game.CheckColor(dropped_number)
	if color != "green"{
		if lengthOfBetsToRed > 0{
			if _, ok := redToBets[color]; ok{
				prize += redToBets[color] * float64(2)
			}
		}
		if lengthOfBetsToBlack > 0{
			if _, ok := blackToBets[color]; ok{
				prize += blackToBets[color] * float64(2)
			}
		}
	}
	return prize, nil
}


//////////////////////////////////////////////////////////////////////////////////////////////
//              NORMALSPINROULETTE (LOOK DOWN)                                  			//
//////////////////////////////////////////////////////////////////////////////////////////////

func (game *GameRoulette) ChooseRandomNumber(array []int) int {
    // Generate a random index within the range of the array length
    randomIndex := rand.Intn(len(array))
    return array[randomIndex]
}


func (game *GameRoulette) GenerateRandomNumberFromArray(numbers []int) int {
    return game.ChooseRandomNumber(numbers)
}

func (game *GameRoulette) GenerateRandomSectorFromArray(number int) string{
	if number >= 1 && number <= 12{
		return "1 st 12"
	}
	if number >= 13 && number <= 24{
		return "2 nd 12"
	}
	if number >= 25 && number <= 36{
		return "3 rd 12"
	}
	return "zero"
}


func (game *GameRoulette) NormalSpinRoulette(redToBets map[string]float64, blackToBets map[string]float64, sectorsToBets map[string]float64, numbersToBets map[int]float64) (float64, error){
	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	lengthOfBetsToRed := len(redToBets)
	lengthOfBetsToBlack := len(blackToBets)

	prize := float64(0)


	dropped_number := game.GenerateRandomNumberFromArray(game.Numbers)
	if lengthOfBetsToNumbers > 0{
		if _, ok := numbersToBets[dropped_number]; ok{
			prize += numbersToBets[dropped_number] * float64(35)
		}
	}
	if lengthOfBetsToSectors > 0{
		dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
		if _, ok := sectorsToBets[dropped_sector]; ok{
			prize += sectorsToBets[dropped_sector] * float64(3)
		}
	}
	color := game.CheckColor(dropped_number)
	if color != "green"{
		if lengthOfBetsToRed > 0{
			if _, ok := redToBets[color]; ok{
				prize += redToBets[color] * float64(2)
			}
		}
		if lengthOfBetsToBlack > 0{
			if _, ok := blackToBets[color]; ok{
				prize += blackToBets[color] * float64(2)
			}
		}
	}

	return prize, nil
}


