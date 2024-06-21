package services

import (
	"fmt"
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

// FIXME
func (game *GameRoulette) CheckNumberBetV2(lengthOfBetsToNumbers int, numbersToBets map[int]float64, dropped_number int) float64{
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
	) (float64, int, error){

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
	fmt.Printf("DROPPED_NUMBER: %d\n", dropped_number)
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

	return prize, dropped_number, nil
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
) (float64, int, error) {
    dropped_number := game.GenerateRandomNumberFromArray(game.Numbers)
    dropped_sector := game.GenerateRandomSectorFromArray(dropped_number)
    prize := 0.0

	// Проверка ставок на конкретные числа
	prize += game.CheckNumberBet(len(numbersToBets), numbersToBets, dropped_number)
	if dropped_number == 0 {
        return prize, dropped_number, nil
    }
    // Проверка ставок на секторы
    prize += game.CheckSectorBet(len(sectorsToBets), sectorsToBets, dropped_sector)
    
    // Проверка ставок на цвет
    prize += game.CheckColorBet(len(blackToBets), len(redToBets), blackToBets, redToBets, dropped_number)
    
    // Проверка ставок на чет/нечет
    prize += game.CheckParityBet(len(evenToBets), len(oddToBets), evenToBets, oddToBets, dropped_number)
    
    // Проверка ставок на 1-18
    prize += game.Check1To18Bet(len(oneToEighteenBets), oneToEighteenBets, dropped_number)
    
    // Проверка ставок на 19-36
    prize += game.Check19To36Bet(len(nineteenToThirtySixBets), nineteenToThirtySixBets, dropped_number)
    
    // Проверка ставок на 1-2-3 2to1
    prize += game.CheckFirst2to1Bet(len(first2To1Bets), first2To1Bets, dropped_number)
    prize += game.CheckSecond2to1Bet(len(second2To1Bets), second2To1Bets, dropped_number)
    prize += game.CheckThird2to1Bet(len(third2To1Bets), third2To1Bets, dropped_number)

    return prize, dropped_number, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////
//              VeryBadSpinRoulette (LOOK DOWN)                                  			//
//////////////////////////////////////////////////////////////////////////////////////////////


func (game *GameRoulette) FindMinBet(
	evenToBets map[string]float64,
	oddToBets map[string]float64,
	redToBets map[string]float64,
	blackToBets map[string]float64,
	sectorsToBets map[string]float64,
	numbersToBets map[[37]int][37]float64, // FIXME
	oneToEighteenBets map[string]float64,
	nineteenToThirtySixBets map[string]float64,
	first2To1Bets map[string]float64,
	second2To1Bets map[string]float64,
	third2To1Bets map[string]float64,) (float64, string){
	
	evenBet := evenToBets["even"]
	oddBet := oddToBets["odd"]
	redBet := redToBets["red"]
	blackBet := blackToBets["black"]
	sectorsBetFirst := sectorsToBets["1 st 12"]
	sectorsBetSecond := sectorsToBets["2 nd 12"]
	sectorsBetThird := sectorsToBets["3 rd 12"]
	// Еще должно быть numbersToBets ТУТ(ключ может быть массивом)
	oneToEighteenBet := oneToEighteenBets["1to18"]
	nineteenToThirtySixBet := nineteenToThirtySixBets["19to36"]
	first2To1Bet := first2To1Bets["2to1"]
	second2To1Bet := second2To1Bets["2to1"]
	third2To1Bet := third2To1Bets["2to1"]
	
	var bets = [13]float64{
		evenBet,
		oddBet,
		redBet,
		blackBet,
		sectorsBetFirst,
		sectorsBetSecond,
		sectorsBetThird,
		sectorsBetThird,
		oneToEighteenBet,
		nineteenToThirtySixBet,
		first2To1Bet,
		second2To1Bet,
		third2To1Bet,
	}
	minBet := bets[0]
	minIndex := 0
	for i := 1; i < 13; i++{
		if bets[i] < minBet{
			minBet = bets[i]
			minIndex = i
		}
	}
	var fields = [13]string{
		"even",
		"odd",
		"red",
		"black",
		"1 st 12",
		"2 nd 12",
		"3 rd 12",
		"1to18",
		"19to36",
		"2to1",
		"2to1",
		"2to1",
	}
	minField := fields[minIndex]
	fmt.Println(minField) 
	return minBet, minField
}


// func (game *GameRoulette) GenerateRandomNumberByBets(
// 	evenToBets map[string]float64,
// 	oddToBets map[string]float64,
// 	redToBets map[string]float64,
// 	blackToBets map[string]float64,
// 	sectorsToBets map[string]float64,
// 	numbersToBets map[[37]int]float64,
// 	oneToEighteenBets map[string]float64,
// 	nineteenToThirtySixBets map[string]float64,
// 	first2To1Bets map[string]float64,
// 	second2To1Bets map[string]float64,
// 	third2To1Bets map[string]float64,
// 	numbers []int) int{
// 	min_bet := game.FindMinBet()	
// }




func (game *GameRoulette) VeryBadSpinRoulette(
	evenToBets map[string]float64,
	oddToBets map[string]float64,
	redToBets map[string]float64,
	blackToBets map[string]float64,
	sectorsToBets map[string]float64,
	numberCell_0 map[int]float64,
	numberCell_1 map[int]float64,
	numberCell_2 map[int]float64,
	numberCell_3 map[int]float64,
	numberCell_4 map[int]float64,
	numberCell_5 map[int]float64,
	numberCell_6 map[int]float64,
	numberCell_7 map[int]float64,
	numberCell_8 map[int]float64,
	numberCell_9 map[int]float64,
	numberCell_10 map[int]float64,
	numberCell_11 map[int]float64,
	numberCell_12 map[int]float64,
	numberCell_13 map[int]float64,
	numberCell_14 map[int]float64,
	numberCell_15 map[int]float64,
	numberCell_16 map[int]float64,
	numberCell_17 map[int]float64,
	numberCell_18 map[int]float64,
	numberCell_19 map[int]float64,
	numberCell_20 map[int]float64,
	numberCell_21 map[int]float64,
	numberCell_22 map[int]float64,
	numberCell_23 map[int]float64,
	numberCell_24 map[int]float64,
	numberCell_25 map[int]float64,
	numberCell_26 map[int]float64,
	numberCell_27 map[int]float64,
	numberCell_28 map[int]float64,
	numberCell_29 map[int]float64,
	numberCell_30 map[int]float64,
	numberCell_31 map[int]float64,
	numberCell_32 map[int]float64,
	numberCell_33 map[int]float64,
	numberCell_34 map[int]float64,
	numberCell_35  map[int]float64,
	numberCell_36 map[int]float64,
	oneToEighteenBets map[string]float64,
	nineteenToThirtySixBets map[string]float64,
	first2To1Bets map[string]float64,
	second2To1Bets map[string]float64,
	third2To1Bets map[string]float64,
	) (float64, int, error){

	lengthOfBetsToSectors := len(sectorsToBets)

	lengthOfBetsNumberCell_0 := len(numberCell_0)
	lengthOfBetsNumberCell_1 := len(numberCell_1)
	lengthOfBetsNumberCell_2 := len(numberCell_2)
	lengthOfBetsNumberCell_3 := len(numberCell_3)
	lengthOfBetsNumberCell_4 := len(numberCell_4)
	lengthOfBetsNumberCell_5 := len(numberCell_5)
	lengthOfBetsNumberCell_6 := len(numberCell_6)
	lengthOfBetsNumberCell_7 := len(numberCell_7)
	lengthOfBetsNumberCell_8 := len(numberCell_8)
	lengthOfBetsNumberCell_9 := len(numberCell_9)
	lengthOfBetsNumberCell_10 := len(numberCell_10)
	lengthOfBetsNumberCell_11 := len(numberCell_11)
	lengthOfBetsNumberCell_12 := len(numberCell_12)
	lengthOfBetsNumberCell_13 := len(numberCell_13)
	lengthOfBetsNumberCell_14 := len(numberCell_14)
	lengthOfBetsNumberCell_15 := len(numberCell_15)
	lengthOfBetsNumberCell_16 := len(numberCell_16)
	lengthOfBetsNumberCell_17 := len(numberCell_17)
	lengthOfBetsNumberCell_18 := len(numberCell_18)
	lengthOfBetsNumberCell_19 := len(numberCell_19)
	lengthOfBetsNumberCell_20 := len(numberCell_20)
	lengthOfBetsNumberCell_21 := len(numberCell_21)
	lengthOfBetsNumberCell_22 := len(numberCell_22)
	lengthOfBetsNumberCell_23 := len(numberCell_23)
	lengthOfBetsNumberCell_24 := len(numberCell_24)
	lengthOfBetsNumberCell_25 := len(numberCell_25)
	lengthOfBetsNumberCell_26 := len(numberCell_26)
	lengthOfBetsNumberCell_27 := len(numberCell_27)
	lengthOfBetsNumberCell_28 := len(numberCell_28)
	lengthOfBetsNumberCell_29 := len(numberCell_29)
	lengthOfBetsNumberCell_30 := len(numberCell_30)
	lengthOfBetsNumberCell_31 := len(numberCell_31)
	lengthOfBetsNumberCell_32 := len(numberCell_32)
	lengthOfBetsNumberCell_33 := len(numberCell_33)
	lengthOfBetsNumberCell_34 := len(numberCell_34)
	lengthOfBetsNumberCell_35 := len(numberCell_35)
	lengthOfBetsNumberCell_36 := len(numberCell_36)
	

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

	
	prize += game.CheckSectorBet(lengthOfBetsToSectors, sectorsToBets, dropped_sector)

	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_0, numberCell_0, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_1, numberCell_1, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_2, numberCell_2, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_3, numberCell_3, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_4, numberCell_4, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_5, numberCell_5, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_6, numberCell_6, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_7, numberCell_7, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_8, numberCell_8, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_9, numberCell_9, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_10, numberCell_10, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_11, numberCell_11, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_12, numberCell_12, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_13, numberCell_13, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_14, numberCell_14, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_15, numberCell_15, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_16, numberCell_16, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_17, numberCell_17, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_18, numberCell_18, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_19, numberCell_19, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_20, numberCell_20, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_21, numberCell_21, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_22, numberCell_22, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_23, numberCell_23, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_24, numberCell_24, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_25, numberCell_25, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_26, numberCell_26, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_27, numberCell_27, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_28, numberCell_28, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_29, numberCell_29, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_30, numberCell_30, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_31, numberCell_31, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_32, numberCell_32, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_33, numberCell_33, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_34, numberCell_34, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_35, numberCell_35, dropped_number)
	prize += game.CheckNumberBetV2(lengthOfBetsNumberCell_36, numberCell_36, dropped_number)



	prize += game.CheckColorBet(lengthOfBetsToBlack, lengthOfBetsToRed, blackToBets, redToBets, dropped_number)
	prize += game.CheckParityBet(lengthOfBetsToEven, lengthOfBetsToOdd, evenToBets, oddToBets, dropped_number)
	prize += game.Check1To18Bet(lengthOfBetsOneToEighteen, oneToEighteenBets, dropped_number)
	prize += game.Check19To36Bet(lengthOfBetsNineteenToThirtySix, nineteenToThirtySixBets, dropped_number)
	prize += game.CheckFirst2to1Bet(lengthOfBetsFirst2To1, first2To1Bets, dropped_number)
	prize += game.CheckSecond2to1Bet(lengthOfBetsSecond2To1, second2To1Bets, dropped_number)
	prize += game.CheckThird2to1Bet(lengthofBetsThird2To1, third2To1Bets, dropped_number)

	return prize, dropped_number, nil
}

