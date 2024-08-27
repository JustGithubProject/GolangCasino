package slots

import (
	"math/rand"
	"time"
)

func CreateWolfGoldPlayingField() [][]int {
	rows, cols := 3, 5
	playingField := make([][]int, rows)
	for i := 0; i < rows; i++ {
		playingField[i] = make([]int, cols)
		for j := 0; j < cols; j++{
			playingField[i][j] = 0
		}
	}
	return playingField
}

type WolfGoldSymbols struct {
	// Low
	jack int // symbol=J order=1
	queen int // symbol=Q order=2
	king int // symbol=K order=3
	ace int // symbol=A order=4

	// Middle
	cougar int // symbol=Cougar order=5
	horse int // symbol=Horse order=6

	// High
	eagle int // symbol=Eagle order=7
	bison int // symbol=Bison order=8
	wolf int // symbol=Wolf order=9

	// Scatter
	scatter int // symbol=Scatter order=10
}


func CheckJackWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.20
	}

	if countAppearance == 4 {
		return 0.80
	}

	if countAppearance == 5 {
		return 2.00
	}

	return 0.0
}


func CheckQueenWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.20
	}

	if countAppearance == 4 {
		return 0.80
	}

	if countAppearance == 5 {
		return 2.00
	}

	return 0.0
}


func CheckKingWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.20
	}

	if countAppearance == 4 {
		return 0.80
	}

	if countAppearance == 5 {
		return 2.00
	}

	return 0.0
}


func CheckAceWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.40
	}

	if countAppearance == 4 {
		return 0.80
	}

	if countAppearance == 5 {
		return 2.00
	}

	return 0.0
}


func CheckCougarWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.40
	}

	if countAppearance == 4 {
		return 2.0
	}

	if countAppearance == 5 {
		return 8.00
	}

	return 0.0
}


func CheckHorseWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.60
	}

	if countAppearance == 4 {
		return 4.00
	}

	if countAppearance == 5 {
		return 12.00
	}

	return 0.0
}


func CheckEagleWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.80
	}

	if countAppearance == 4 {
		return 6.00
	}

	if countAppearance == 5 {
		return 16.00
	}

	return 0.0
}


func CheckBisonWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.00
	}

	if countAppearance == 4 {
		return 10.00
	}

	if countAppearance == 5 {
		return 20.00
	}

	return 0.0
}


func CheckWolfWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.00
	}

	if countAppearance == 4 {
		return 10.00
	}

	if countAppearance == 5 {
		return 20.00
	}

	return 0.0
}


func CheckScatterWolfGoldSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.00
	}

	return 0.0
}


func CalculateWolfGoldPaymentsNormalMode(playingField [][]int, bet float64, balance float64) float64 {

	// Checking the payout for each set of symbols
	payoutJackX := CheckJackWolfGoldSymbolPlayed(playingField, 1)
	payoutQueenX := CheckQueenWolfGoldSymbolPlayed(playingField, 2)
	payoutKingX := CheckKingWolfGoldSymbolPlayed(playingField, 3)
	payoutAceX := CheckAceWolfGoldSymbolPlayed(playingField, 4)
	payoutEagleX := CheckEagleWolfGoldSymbolPlayed(playingField, 5)
	payoutBisonX := CheckBisonWolfGoldSymbolPlayed(playingField, 6)
	payoutWolfX := CheckWolfWolfGoldSymbolPlayed(playingField, 7)

	// Calculating total payout
	totalPayout := 0.0

	if payoutJackX > 0.0 {
		totalPayout += bet * payoutJackX
	}

	if payoutQueenX > 0.0 {
		totalPayout += bet * payoutQueenX
	}

	if payoutKingX > 0.0 {
		totalPayout += bet * payoutKingX
	}

	if payoutAceX > 0.0 {
		totalPayout += bet * payoutAceX
	}

	if payoutEagleX > 0.0 {
		totalPayout += bet * payoutEagleX
	}

	if payoutBisonX > 0.0 {
		totalPayout += bet * payoutBisonX
	}

	if payoutWolfX > 0.0 {
		totalPayout += bet * payoutWolfX
	}


	// New balance = Initial balance − Bet + Win
	balance = balance - bet + totalPayout
	return balance
}


func WeightedWolfGoldRandomChoice(values []int, weights []float64) int {
	if len(values) != len(weights) {
		panic("Values and weights must be the same length")
	}

	// Weight to invert
	invWeights := make([]float64, len(weights))
	for i, w := range weights {
		invWeights[i] = 1.0 / w
	}

	// Make a list of cumulative sums
	cumSum := make([]float64, len(invWeights))
	cumSum[0] = invWeights[0]
	for i := 1; i < len(invWeights); i++ {
		cumSum[i] = cumSum[i-1] + invWeights[i]
	}

	// Generate a random number
	r := rand.Float64() * cumSum[len(cumSum)-1]

	// Находим элемент, соответствующий случайному числу
	for i, cs := range cumSum {
		if r < cs {
			return values[i]
		}
	}

	return values[len(values)-1]
}

func GenerateWolfGoldRandomNumberNormalMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights := []float64{
		5.0, 5.0, 10.0, 15.0, 20.0, 25.0, 50.0, 55.0, 60.0, 65.0, 100.0,
	}
	randomNumber := WeightedDogHouseRandomChoice(values, weights)
	return randomNumber
}

func GenerateWolfGoldRandomNumberBonusMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights := []float64{
		5.0, 5.0, 5.0, 5.0, 15.0, 15.0, 20.0, 350.0, 400.0, 450.0, 500.0,
	}
	randomNumber := WeightedDogHouseRandomChoice(values, weights)
	return randomNumber
}

func init() {
	rand.Seed(time.Now().UnixNano())
}


func GenerateWolfGoldPlayingFieldNormalMode() [][]int {
	cols, rows := 3, 5
	playingField := CreateDogHousePlayingField()
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			playingField[i][j] = GenerateDogHouseRandomNumberNormalMode()
		}
	}
	return playingField
}

func GenerateWolfGoldPlayingFieldBonusMode() [][]int {
	playingField := CreatePlayingField()
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			playingField[i][j] = GenerateDogHouseRandomNumberBonusMode()
		}
	}
	return playingField
}