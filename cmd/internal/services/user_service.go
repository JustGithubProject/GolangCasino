package services

import "fmt"

type UserPlayer struct {
	TypeOfGame GameRoulette
	Balance    float64
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
) (float64, error) {
	fmt.Println("Зашел в Normal Play")
	// Падает
	numbers := InitNumbersArray()
	user.TypeOfGame.Numbers = numbers
	fmt.Println(user.TypeOfGame.Numbers)
	prize, err := user.TypeOfGame.NormalSpinRoulette(
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
		return 0, &GameError{Message: "Game play error: " + err.Error()}
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
	fmt.Printf("PRIZE=%.2f\n", prize)
	fmt.Printf("totalBet=%.2f\n", prize)
	fmt.Printf("UserBalance: %.2f\n", user.Balance)
	user.updateBalance(prize, totalBet)
	fmt.Printf("After update UserBalance: %.2f\n", user.Balance)
	return user.Balance, nil
}

func (user *UserPlayer) updateBalance(prize float64, bet float64) {
	if prize > bet {
		user.Balance += prize
		if prize == bet {
			user.Balance += 0
		}
	} else {
		user.Balance -= bet
	}
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
) (float64, error) {
	InitWeights(user.TypeOfGame.WeightsForNumbers, 37)
	ShuffleWeights(user.TypeOfGame.WeightsForNumbers)
	numbers := InitNumbersArray()
	user.TypeOfGame.Numbers = numbers

	prize, err := user.TypeOfGame.UnfairSpinRoulette(
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
		return 0, &GameError{Message: "Game play error: " + err.Error()}
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
	return user.Balance, nil
}
