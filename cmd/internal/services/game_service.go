package services


import (
    "math/rand"
)


type GameRoulette struct{
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

func (game *GameRoulette) IsEvenOrOdd(number int) string {
	if number % 2 == 0{
		return "even"
	}
	return "odd"
}

func (game *GameRoulette) CheckNumberBet(lengthOfBetsToNumbers int, numbersToBets map[int]float64, dropped_number int) float64{
	if lengthOfBetsToNumbers > 0{
		if _, ok := numbersToBets[dropped_number]; ok{
			return numbersToBets[dropped_number] * float64(35)
		}
	}
	return float64(0)
}

func (game *GameRoulette) CheckSectorBet(lengthOfBetsToSectors int, sectorsToBets map[string]float64, dropped_sector string) float64{
	if lengthOfBetsToSectors > 0{
		if _, ok := sectorsToBets[dropped_sector]; ok{
			return sectorsToBets[dropped_sector] * float64(3)
		}
	}
	return float64(0)
}

func (game *GameRoulette) CheckColorBet(lengthOfBetsToBlack int, lengthOfBetsToRed int, blackToBets map[string]float64, redToBets map[string]float64, dropped_number int) float64{
	color := game.CheckColor(dropped_number)
	if color != "green"{
		if lengthOfBetsToRed > 0{
			if _, ok := redToBets[color]; ok{
				return redToBets[color] * float64(2)
			}
		}
		if lengthOfBetsToBlack > 0{
			if _, ok := blackToBets[color]; ok{
				return blackToBets[color] * float64(2)
			}
		}
	}
	return float64(0)
}


func (game *GameRoulette) CheckParityBet(
	lengthOfBetsToEven int,
	lengthOfBetsToOdd int,
	evenToBets map[string]float64,
	oddToBets map[string]float64,
	dropped_number int,
	) float64{
	parity := game.IsEvenOrOdd(dropped_number)
	if lengthOfBetsToEven > 0{
		if _, ok := evenToBets[parity]; ok{
			return evenToBets[parity] * float64(2)
		}
	}
	if lengthOfBetsToOdd > 0{
		if _, ok := oddToBets[parity]; ok{
			return oddToBets[parity] * float64(2)
		}
	}
	return float64(0)
}

func (game *GameRoulette) Check1To18Bet(
	lengthOfBetsOneToEighteen int,
	oneToEighteenBets map[string]float64,
	dropped_number int,
	) float64{
	if lengthOfBetsOneToEighteen > 0{
		for i := 1; i <= 18; i++{
			if dropped_number == i{
				return oneToEighteenBets["1to18"] * float64(2)
			}
		}
	}
	return float64(0)
}

func (game *GameRoulette) Check19To36Bet(
	lengthOfBetsNineteenToThirtySix int,
	nineteenToThirtySixBets map[string]float64,
	dropped_number int,
	) float64{
	if lengthOfBetsNineteenToThirtySix > 0{
		for i := 19; i <= 36; i++{
			if dropped_number == i{
				return nineteenToThirtySixBets["19to36"] * float64(2)
			}
		}
	}

	return float64(0)
}


func (game *GameRoulette) CheckFirst2to1Bet(
	lengthOfBetsFirst2to1 int,
	first2to1Bets map[string]float64,
	dropped_number int,
) float64{
	// Инициализация массива 1, 4, 7 ... первый подсектор 
	var first2To1Array[12]int
	j := 1
	for i := 0; i < 12; i++{
		first2To1Array[i] = j
		j += 3
	}
	if lengthOfBetsFirst2to1 > 0{
		for i := 0; i < 12; i++{
			if dropped_number == first2To1Array[i]{
				return first2to1Bets["2to1"] * float64(3)
			}
		}
	}

	return 0.0

}

func (game *GameRoulette) CheckSecond2to1Bet(
	lengthOfBetsSecond2to1 int,
	second2to1Bets map[string]float64,
	dropped_number int,
) float64{
	var second2To1Array[12]int
	// Инициализация массива 2, 5, 8 ... второй подсектор
	j := 2
	for i := 0; i < 12; i++{
		second2To1Array[i] = j
		j += 3
	}

	if lengthOfBetsSecond2to1 > 0{
		for i := 0; i < 12; i++{
			if dropped_number == second2To1Array[i]{
				return second2to1Bets["2to1"] * float64(3)
			}
		}
	}
	return 0.0

	
}

func (game *GameRoulette) CheckThird2to1Bet(
	lengthOfBetsThird2to1 int,
	third2to1Bets map[string]float64,
	dropped_number int,
) float64{
	var third2To1Array[12]int
	// Инициализация массива 3, 6, 9 ... третий подсектор
	j := 3
	for i := 0; i < 12; i++{
		third2To1Array[i] = j
		j += 3
	}

	if lengthOfBetsThird2to1 > 0{
		for i := 0; i < 12; i++{
			if dropped_number == third2To1Array[i]{
				return third2to1Bets["2to1"] * float64(3)
			}
		}
	}
	return 0.0
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


func (game *GameRoulette) UnfairSpinRoulette(
	evenToBets map[string]float64,
	oddToBets map[string]float64,
	redToBets map[string]float64,
	blackToBets map[string]float64,
	sectorsToBets map[string]float64,
	numbersToBets map[int]float64,
	oneToEighteenBets map[string]float64,
	nineteenToThirtySixBets map[string]float64,
	first2To1Bets map[string]float64,
	second2To1Bets map[string]float64,
	third2To1Bets map[string]float64,
	) (float64, error){

	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	lengthOfBetsToRed := len(redToBets)
	lengthOfBetsToBlack := len(blackToBets)
	lengthOfBetsToEven := len(evenToBets)
	lengthOfBetsToOdd := len(oddToBets)
	lengthOfBetsOneToEighteen := len(oneToEighteenBets)
	lengthOfBetsNineteenToThirtySix := len(nineteenToThirtySixBets)
	lengthOfBetsFirst2To1 := len(first2To1Bets)
	lengthOfBetsSecond2To1 := len(second2To1Bets)
	lengthofBetsThird2To1 := len(third2To1Bets)


	dropped_number := game.GenerateRandomNumberByWeight(game.Numbers, game.WeightsForNumbers)
	dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
	prize := 0.0

	prize += game.CheckNumberBet(lengthOfBetsToNumbers, numbersToBets, dropped_number)
	prize += game.CheckSectorBet(lengthOfBetsToSectors, sectorsToBets, dropped_sector)
	prize += game.CheckColorBet(lengthOfBetsToBlack, lengthOfBetsToRed, blackToBets, redToBets, dropped_number)
	prize += game.CheckParityBet(lengthOfBetsToEven, lengthOfBetsToOdd, evenToBets, oddToBets, dropped_number)
	prize += game.Check1To18Bet(lengthOfBetsOneToEighteen, oneToEighteenBets, dropped_number)
	prize += game.Check19To36Bet(lengthOfBetsNineteenToThirtySix, nineteenToThirtySixBets, dropped_number)
	prize += game.CheckFirst2to1Bet(lengthOfBetsFirst2To1, first2To1Bets, dropped_number)
	prize += game.CheckSecond2to1Bet(lengthOfBetsSecond2To1, second2To1Bets, dropped_number)
	prize += game.CheckThird2to1Bet(lengthofBetsThird2To1, third2To1Bets, dropped_number)

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


func (game *GameRoulette) NormalSpinRoulette(
	evenToBets map[string]float64,
	oddToBets map[string]float64,
	redToBets map[string]float64,
	blackToBets map[string]float64,
	sectorsToBets map[string]float64,
	numbersToBets map[int]float64,
	oneToEighteenBets map[string]float64,
	nineteenToThirtySixBets map[string]float64,
	first2To1Bets map[string]float64,
	second2To1Bets map[string]float64,
	third2To1Bets map[string]float64,
	) (float64, error){

	lengthOfBetsToSectors := len(sectorsToBets)
	lengthOfBetsToNumbers := len(numbersToBets)
	lengthOfBetsToRed := len(redToBets)
	lengthOfBetsToBlack := len(blackToBets)
	lengthOfBetsToEven := len(evenToBets)
	lengthOfBetsToOdd := len(oddToBets)
	lengthOfBetsOneToEighteen := len(oneToEighteenBets)
	lengthOfBetsNineteenToThirtySix := len(nineteenToThirtySixBets)
	lengthOfBetsFirst2To1 := len(first2To1Bets)
	lengthOfBetsSecond2To1 := len(second2To1Bets)
	lengthofBetsThird2To1 := len(third2To1Bets)

	dropped_number := game.GenerateRandomNumberFromArray(game.Numbers)
	dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
	prize := 0.0

	prize += game.CheckNumberBet(lengthOfBetsToNumbers, numbersToBets, dropped_number)
	prize += game.CheckSectorBet(lengthOfBetsToSectors, sectorsToBets, dropped_sector)
	prize += game.CheckColorBet(lengthOfBetsToBlack, lengthOfBetsToRed, blackToBets, redToBets, dropped_number)
	prize += game.CheckParityBet(lengthOfBetsToEven, lengthOfBetsToOdd, evenToBets, oddToBets, dropped_number)
	prize += game.Check1To18Bet(lengthOfBetsOneToEighteen, oneToEighteenBets, dropped_number)
	prize += game.Check19To36Bet(lengthOfBetsNineteenToThirtySix, nineteenToThirtySixBets, dropped_number)
	prize += game.CheckFirst2to1Bet(lengthOfBetsFirst2To1, first2To1Bets, dropped_number)
	prize += game.CheckSecond2to1Bet(lengthOfBetsSecond2To1, second2To1Bets, dropped_number)
	prize += game.CheckThird2to1Bet(lengthofBetsThird2To1, third2To1Bets, dropped_number)

	return prize, nil
}


