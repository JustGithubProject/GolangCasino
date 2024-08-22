package slots

import (
	"math/rand"
	"time"
)

func CreateDogHousePlayingField() [][]int {
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

type DogHouseSymbols struct {
	// Low value symbols
	ten   int // symbol=10 order=1
	jack  int // symbol=J order=2
	queen int // symbol=Q order=3
	king  int // symbol=K order=4
	ace   int // symbol=A order=5

	// Medium value symbols
	bone   int // symbol=Bone order=6
	collar int // symbol=Collar order=7

	// High Value symbols (Dogs)
	dachshund int // symbol=Dachshund order=8
	pug       int // symbol=Pug order=9
	spitz     int // symbol=Spitz order=10
	boxer     int // symbol=Boxer order=11

	// Bonus
	paws int // symbol=Paws order=12
}

/*
	BLOCK: Checking multipliers
*/

func CountMatrixElement(playingField [][]int, symbol int) int {
	countAppearance := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}
	return countAppearance
}

func CheckTenSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.10
	}

	if countAppearance == 4 {
		return 0.25
	}

	if countAppearance == 5 {
		return 1.25
	}

	return 0.0
}

func CheckJackSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.10
	}

	if countAppearance == 4 {
		return 0.25
	}

	if countAppearance == 5 {
		return 1.25
	}

	return 0.0
}

func CheckQueenSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.10
	}

	if countAppearance == 4 {
		return 0.25
	}

	if countAppearance == 5 {
		return 1.25
	}

	return 0.0
}

func CheckKingSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.25
	}

	if countAppearance == 4 {
		return 0.50
	}

	if countAppearance == 5 {
		return 2.5
	}

	return 0.0
}

func CheckAceSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.25
	}

	if countAppearance == 4 {
		return 0.50
	}

	if countAppearance == 5 {
		return 2.5
	}

	return 0.0
}

func CheckBoneSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.40
	}

	if countAppearance == 4 {
		return 1.00
	}

	if countAppearance == 5 {
		return 5.00
	}

	return 0.0
}

func CheckCollarSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 0.60
	}

	if countAppearance == 4 {
		return 1.25
	}

	if countAppearance == 5 {
		return 7.5
	}

	return 0.0
}

func CheckDachshundSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.00
	}

	if countAppearance == 4 {
		return 2.00
	}

	if countAppearance == 5 {
		return 10.00
	}

	return 0.0
}

func CheckPugSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.25
	}

	if countAppearance == 4 {
		return 3.00
	}

	if countAppearance == 5 {
		return 15.00
	}

	return 0.0
}

func CheckSpitzSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 1.75
	}

	if countAppearance == 4 {
		return 5.00
	}

	if countAppearance == 5 {
		return 25.00
	}

	return 0.0
}

func CheckBoxerSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3 {
		return 2.50
	}

	if countAppearance == 4 {
		return 7.50
	}

	if countAppearance == 5 {
		return 37.50
	}

	return 0.0
}

func CheckPawsSymbolPlayed(playingField [][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)
	if countAppearance == 3 {
		// TODO: ...
	}

	if countAppearance == 4 {
		// TODO: ...
	}

	if countAppearance == 5 {
		// TODO: ...
	}

	return 0.0
}

func CalculateDogHousePaymentsNormalMode(playingField [][]int, bet float64, balance float64) float64 {

	// Checking the payout for each set of symbols
	payoutTenX := CheckTenSymbolPlayed(playingField, 1)
	payoutJackX := CheckJackSymbolPlayed(playingField, 2)
	payoutQueenX := CheckQueenSymbolPlayed(playingField, 3)
	payoutKingX := CheckKingSymbolPlayed(playingField, 4)
	payoutAceX := CheckAceSymbolPlayed(playingField, 5)
	payoutBoneX := CheckBoneSymbolPlayed(playingField, 6)
	payoutCollarX := CheckCollarSymbolPlayed(playingField, 7)
	payoutDachshundX := CheckDachshundSymbolPlayed(playingField, 8)
	payoutPugX := CheckPugSymbolPlayed(playingField, 9)
	payoutSpitzX := CheckSpitzSymbolPlayed(playingField, 10)
	payoutBoxerX := CheckBoxerSymbolPlayed(playingField, 11)

	// Calculating total payout
	totalPayout := 0.0

	if payoutTenX > 0.0 {
		totalPayout += bet * payoutTenX
	}

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

	if payoutBoneX > 0.0 {
		totalPayout += bet * payoutBoneX
	}

	if payoutCollarX > 0.0 {
		totalPayout += bet * payoutCollarX
	}

	if payoutDachshundX > 0.0 {
		totalPayout += bet * payoutDachshundX
	}

	if payoutPugX > 0.0 {
		totalPayout += bet * payoutPugX
	}

	if payoutSpitzX > 0.0 {
		totalPayout += bet * payoutSpitzX
	}

	if payoutBoxerX > 0.0 {
		totalPayout += bet * payoutBoxerX
	}

	// New balance = Initial balance − Bet + Win
	balance = balance - bet + totalPayout
	return balance
}

func WeightedDogHouseRandomChoice(values []int, weights []float64) int {
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

func GenerateDogHouseRandomNumberNormalMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	weights := []float64{
		5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 15.0, 350.0, 400.0, 450.0, 500.0,
	}
	randomNumber := WeightedDogHouseRandomChoice(values, weights)
	return randomNumber
}

func GenerateDogHouseRandomNumberBonusMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	weights := []float64{
		10.0, 50.0, 100.0, 150.0, 200.0, 250.0, 300.0, 350.0, 450.0, 500.0,
		550.0, 600.0, 650.0, 700.0, 750.0, 10000.0, 100000.0,
	}
	randomNumber := WeightedDogHouseRandomChoice(values, weights)
	return randomNumber
}

func init() {
	rand.Seed(time.Now().UnixNano())
}


func GenerateDogHousePlayingFieldNormalMode() [][]int {
	cols, rows := 3, 5
	playingField := CreateDogHousePlayingField()
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			playingField[i][j] = GenerateDogHouseRandomNumberNormalMode()
		}
	}
	return playingField
}

func GenerateDogHousePlayingFieldBonusMode() [][]int {
	playingField := CreatePlayingField()
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			playingField[i][j] = GenerateDogHouseRandomNumberBonusMode()
		}
	}
	return playingField
}

func DogHouseSpin(isBonusMode bool, bet float64, balance float64) ([][]int, float64) {
	if isBonusMode {
		currentPlayingField := GenerateDogHousePlayingFieldBonusMode()
		currentBalance := CalculateDogHousePaymentsNormalMode(currentPlayingField, bet, balance) // Isn't correct.
		return currentPlayingField, currentBalance
	}
	currentPlayingField := GenerateDogHousePlayingFieldNormalMode()
	currentBalance := CalculateDogHousePaymentsNormalMode(currentPlayingField, bet, balance)
	return currentPlayingField, currentBalance
}