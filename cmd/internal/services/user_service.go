package services



type UserPlayer struct {
	TypeOfGame GameRoulette
	Balance		float64
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
	InitNumbersArray(user.TypeOfGame.Numbers)
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
	if err != nil{
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


func (user *UserPlayer) updateBalance(money float64, bet float64){
	if money > bet{
		user.Balance += money
	}else{
		user.Balance -= money
	}
}

func (user *UserPlayer) sumMapValues(bets map[string]float64) float64 {
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
	total += user.sumMapValues(numbersToBets)
	total += user.sumMapValues(oneToEighteenBets)
	total += user.sumMapValues(nineteenToThirtySixBets)
	total += user.sumMapValues(first2To1Bets)
	total += user.sumMapValues(second2To1Bets)
	total += user.sumMapValues(third2To1Bets)

	return total
}

func (user *UserPlayer) UnFairPlay(guess_sector string, guess_number int, bet float64, gameName string) (float64, error){
	user.TypeOfGame.GameName = gameName
	InitWeights(user.TypeOfGame.WeightsForNumbers, 37)
	InitWeights(user.TypeOfGame.WeightsForSectors, len(user.TypeOfGame.Sectors))


	ShuffleWeights(user.TypeOfGame.WeightsForNumbers)
	ShuffleWeights(user.TypeOfGame.WeightsForSectors)


	InitNumbersArray(user.TypeOfGame.Numbers)
	money, err := user.TypeOfGame.UnfairSpinRoulette(guess_sector, guess_number, bet)
	if err != nil{
		return 0, &GameError{Message: "Game play error: " + err.Error()}
	}
	user.updateBalance(money, bet)
	return user.Balance, nil
}


