package services

import "fmt"

type UserPlayer struct {
	TypeOfGame GameRoulette
	Balance    float64
}

func (user *UserPlayer) getTotalBetV2(
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
	third2To1Bets map[string]float64) float64 {

	var total float64

	// Добавляем суммы ставок из карты ставок для каждого типа
	total += user.sumMapValues(evenToBets)
	total += user.sumMapValues(oddToBets)
	total += user.sumMapValues(redToBets)
	total += user.sumMapValues(blackToBets)
	total += user.sumMapValues(sectorsToBets)
	total += user.sumMapValuesForNumbers(numberCell_0)
	total += user.sumMapValuesForNumbers(numberCell_1)
	total += user.sumMapValuesForNumbers(numberCell_2)
	total += user.sumMapValuesForNumbers(numberCell_3)
	total += user.sumMapValuesForNumbers(numberCell_4)
	total += user.sumMapValuesForNumbers(numberCell_5)
	total += user.sumMapValuesForNumbers(numberCell_6)
	total += user.sumMapValuesForNumbers(numberCell_7)
	total += user.sumMapValuesForNumbers(numberCell_8)
	total += user.sumMapValuesForNumbers(numberCell_9)
	total += user.sumMapValuesForNumbers(numberCell_10)
	total += user.sumMapValuesForNumbers(numberCell_11)
	total += user.sumMapValuesForNumbers(numberCell_12)
	total += user.sumMapValuesForNumbers(numberCell_13)
	total += user.sumMapValuesForNumbers(numberCell_14)
	total += user.sumMapValuesForNumbers(numberCell_15)
	total += user.sumMapValuesForNumbers(numberCell_16)
	total += user.sumMapValuesForNumbers(numberCell_17)
	total += user.sumMapValuesForNumbers(numberCell_18)
	total += user.sumMapValuesForNumbers(numberCell_19)
	total += user.sumMapValuesForNumbers(numberCell_20)
	total += user.sumMapValuesForNumbers(numberCell_21)
	total += user.sumMapValuesForNumbers(numberCell_22)
	total += user.sumMapValuesForNumbers(numberCell_23)
	total += user.sumMapValuesForNumbers(numberCell_24)
	total += user.sumMapValuesForNumbers(numberCell_25)
	total += user.sumMapValuesForNumbers(numberCell_26)
	total += user.sumMapValuesForNumbers(numberCell_27)
	total += user.sumMapValuesForNumbers(numberCell_28)
	total += user.sumMapValuesForNumbers(numberCell_29)
	total += user.sumMapValuesForNumbers(numberCell_30)
	total += user.sumMapValuesForNumbers(numberCell_31)
	total += user.sumMapValuesForNumbers(numberCell_32)
	total += user.sumMapValuesForNumbers(numberCell_33)
	total += user.sumMapValuesForNumbers(numberCell_34)
	total += user.sumMapValuesForNumbers(numberCell_35)
	total += user.sumMapValuesForNumbers(numberCell_36)
	total += user.sumMapValues(oneToEighteenBets)
	total += user.sumMapValues(nineteenToThirtySixBets)
	total += user.sumMapValues(first2To1Bets)
	total += user.sumMapValues(second2To1Bets)
	total += user.sumMapValues(third2To1Bets)

	return total
}

func (user *UserPlayer) VeryBadPlay(
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
	fmt.Println("Зашел в VeryBadPlay")
	numbers := InitNumbersArray()
	user.TypeOfGame.Numbers = numbers
	fmt.Println(user.TypeOfGame.Numbers)
	prize, dropped_number, err := user.TypeOfGame.VeryBadSpinRoulette(
		evenToBets,
		oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numberCell_0,
		numberCell_1,
		numberCell_2,
		numberCell_3,
		numberCell_4,
		numberCell_5,
		numberCell_6,
		numberCell_7,
		numberCell_8,
		numberCell_9,
		numberCell_10,
		numberCell_11,
		numberCell_12,
		numberCell_13,
		numberCell_14,
		numberCell_15,
		numberCell_16,
		numberCell_17,
		numberCell_18,
		numberCell_19,
		numberCell_20,
		numberCell_21,
		numberCell_22,
		numberCell_23,
		numberCell_24,
		numberCell_25,
		numberCell_26,
		numberCell_27,
		numberCell_28,
		numberCell_29,
		numberCell_30,
		numberCell_31,
		numberCell_32,
		numberCell_33,
		numberCell_34,
		numberCell_35,
		numberCell_36,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	if err != nil {
		return 0, -1, &GameError{Message: "Game play error: " + err.Error()}
	}
	totalBet := user.getTotalBetV2(
		evenToBets,
		oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numberCell_0,
		numberCell_1,
		numberCell_2,
		numberCell_3,
		numberCell_4,
		numberCell_5,
		numberCell_6,
		numberCell_7,
		numberCell_8,
		numberCell_9,
		numberCell_10,
		numberCell_11,
		numberCell_12,
		numberCell_13,
		numberCell_14,
		numberCell_15,
		numberCell_16,
		numberCell_17,
		numberCell_18,
		numberCell_19,
		numberCell_20,
		numberCell_21,
		numberCell_22,
		numberCell_23,
		numberCell_24,
		numberCell_25,
		numberCell_26,
		numberCell_27,
		numberCell_28,
		numberCell_29,
		numberCell_30,
		numberCell_31,
		numberCell_32,
		numberCell_33,
		numberCell_34,
		numberCell_35,
		numberCell_36,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	
	user.updateBalance(prize, totalBet)
	return user.Balance, dropped_number, nil
}



func (user *UserPlayer) NormalPlay(
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
	fmt.Println("Зашел в Normal Play")
	numbers := InitNumbersArray()
	user.TypeOfGame.Numbers = numbers
	fmt.Println(user.TypeOfGame.Numbers)
	prize, dropped_number, err := user.TypeOfGame.NormalSpinRoulette(
		evenToBets,
		oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numbersToBets,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	if err != nil {
		return 0, -1, &GameError{Message: "Game play error: " + err.Error()}
	}
	totalBet := user.getTotalBet(evenToBets, oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numbersToBets,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	
	user.updateBalance(prize, totalBet)
	return user.Balance, dropped_number, nil
}

func (user *UserPlayer) updateBalance(prize float64, bet float64) {
    user.Balance += prize - bet
}

func (user *UserPlayer) sumMapValues(bets map[string]float64) float64 {
	sum := 0.0
	for _, value := range bets {
		sum += value
	}
	return sum
}

func (user *UserPlayer) sumMapValuesForNumbers(bets map[int]float64) float64 {
	sum := 0.0
	for _, value := range bets {
		sum += value
	}
	return sum
}


func (user *UserPlayer) sumMapValuesForNumbersV2(bets map[[37]int][37]float64) float64 {
	sum := 0.0

	for _, value := range bets{
		for i := 0; i < 37; i++{
			sum += value[i]
		}
	}
	return sum
}

func (user *UserPlayer) getTotalBet(evenToBets map[string]float64, oddToBets map[string]float64,
	redToBets map[string]float64, blackToBets map[string]float64,
	sectorsToBets map[string]float64, numbersToBets map[int]float64,
	oneToEighteenBets map[string]float64, nineteenToThirtySixBets map[string]float64,
	first2To1Bets map[string]float64, second2To1Bets map[string]float64,
	third2To1Bets map[string]float64) float64 {

	var total float64

	// Добавляем суммы ставок из карты ставок для каждого типа
	total += user.sumMapValues(evenToBets)
	total += user.sumMapValues(oddToBets)
	total += user.sumMapValues(redToBets)
	total += user.sumMapValues(blackToBets)
	total += user.sumMapValues(sectorsToBets)
	total += user.sumMapValuesForNumbers(numbersToBets)
	total += user.sumMapValues(oneToEighteenBets)
	total += user.sumMapValues(nineteenToThirtySixBets)
	total += user.sumMapValues(first2To1Bets)
	total += user.sumMapValues(second2To1Bets)
	total += user.sumMapValues(third2To1Bets)

	return total
}

func (user *UserPlayer) UnFairPlay(
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
	weights_arr := InitWeights(37)
	user.TypeOfGame.WeightsForNumbers = weights_arr
	ShuffleWeights(user.TypeOfGame.WeightsForNumbers)
	numbers := InitNumbersArray()
	user.TypeOfGame.Numbers = numbers

	prize, dropped_number, err := user.TypeOfGame.UnfairSpinRoulette(
		evenToBets,
		oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numbersToBets,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	if err != nil {
		return 0, -1, &GameError{Message: "Game play error: " + err.Error()}
	}
	totalBet := user.getTotalBet(evenToBets, oddToBets,
		redToBets,
		blackToBets,
		sectorsToBets,
		numbersToBets,
		oneToEighteenBets,
		nineteenToThirtySixBets,
		first2To1Bets,
		second2To1Bets,
		third2To1Bets,
	)
	fmt.Printf("User balance: %.2f\n", user.Balance)
	fmt.Printf("UnfairPrize: %.2f\n", prize)
	user.updateBalance(prize, totalBet)
	fmt.Printf("UnfairBalance after update: %.2f\n", user.Balance)
	return user.Balance, dropped_number, nil
}
