package services



type UserPlayer struct {
	TypeOfGame GameRoulette
	Balance		float64
}

func (user *UserPlayer) NormalPlay(guess_sectors []string, guess_numbers []int, bet float64, gameName string) (float64, error) {
	user.TypeOfGame.GameName = gameName

	InitNumbersArray(user.TypeOfGame.Numbers)
	money, err := user.TypeOfGame.NormalSpinRoulette(guess_sectors, guess_numbers, bet)
	if err != nil{
		return 0, &GameError{Message: "Game play error: " + err.Error()}
	}
	user.updateBalance(money, bet)
	return user.Balance, nil
}


func (user *UserPlayer) updateBalance(money float64, bet float64){
	if money > bet{
		user.Balance += money
	}else{
		user.Balance -= money
	}
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


