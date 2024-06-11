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
	numbersToBets map[[37]int]float64,
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
	total += user.sumMapValuesForNumbersV2(numbersToBets)
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
	numbersToBets map[[37]int]float64,
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
	totalBet := user.getTotalBetV2(
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
