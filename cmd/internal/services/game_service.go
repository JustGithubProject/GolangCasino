package services


import (
    "math/rand"
)


type GameRoulette struct{
	GameName string
	Numbers []int
	Sectors []string

    WeightsForNumbers []int
    WeightsForSectors []int
}

//////////////////////////////////////////////////////////////////////////////////////////////
//              METHODS FOR UNFAIRSPINROULETTE   (LOOK DOWN)                                //
//////////////////////////////////////////////////////////////////////////////////////////////
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


func (game *GameRoulette) UnfairSpinRoulette(sectorsToBets map[string]float64, numbersToBets map[int]float64) (float64, error){
	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	dropped_number := game.GenerateRandomNumberByWeight(game.Numbers, game.WeightsForNumbers)
	prize := float64(0)
	
	if lengthOfBetsToNumbers > 0{
		if _, ok := numbersToBets[dropped_number]; ok{
			intermediateValue := numbersToBets[dropped_number] * float64(35)
			prize += intermediateValue
		}
	}
	if lengthOfBetsToSectors > 0{
		dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
		if _, ok := sectorsToBets[dropped_sector]; ok{
			intermediateValue := sectorsToBets[dropped_sector] * float64(3)
			prize += intermediateValue
		}
	}

	return prize, nil
}


//////////////////////////////////////////////////////////////////////////////////////////////
//              METHODS FOR NORMALSPINROULETTE (LOOK DOWN)                                  //
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


func (game *GameRoulette) NormalSpinRoulette(sectorsToBets map[string]float64, numbersToBets map[int]float64) (float64, error){
	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	prize := float64(0)
	dropped_number := game.GenerateRandomNumberFromArray(game.Numbers)
	if lengthOfBetsToNumbers > 0{
		if _, ok := numbersToBets[dropped_number]; ok{
			intermediateValue := numbersToBets[dropped_number] * float64(35)
			prize += intermediateValue
		}
	}
	if lengthOfBetsToSectors > 0{
		dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
		if _, ok := sectorsToBets[dropped_sector]; ok{
			intermediateValue := sectorsToBets[dropped_sector] * float64(3)
			prize += intermediateValue
		}
	}
	return prize, nil
}


