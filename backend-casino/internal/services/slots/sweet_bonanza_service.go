package slots

import (
	"math/rand"
	"time"
)

// Each game element will be an integer
type GameSymbols struct {
	Banana int // 1
	Grapes int // 2
	Watermelon int // 3
	Plum int // 4
	Apple int // 5
	BlueCandy int // 6
	GreenCandy int // 7
	PurpleCandy int // 8
	RedCandy int // 9
	Scatter int // 10

	// It will appear in the free spin
	Bomb2X int // 11
	Bomb3X int // 12
	Bomb5X int // 13
	Bomb10X int // 14
	Bomb25X int // 15
	Bomb50X int // 16
	Bomb100X int // 17
}


func CreatePlayingField() [][]int{
	rows, cols := 5, 6
	playingField := make([][]int, rows)
    for i := range playingField {
        playingField[i] = make([]int, cols)
        for j := range playingField[i] {
            playingField[i][j] = 0
        }
    }
	return playingField
}




func CheckBananasPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 0.25
	}
	if countAppearance > 9 && countAppearance < 12{
		return 0.75
	}
	if countAppearance >= 12{
		return 2.0
	}
	return 0.0
}

func CheckGrapesPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 0.40
	}
	if countAppearance > 9 && countAppearance < 12{
		return 0.90
	}
	if countAppearance >= 12{
		return 4.0
	}
	return 0.0
}


func CheckWatermelonPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 0.50
	}
	if countAppearance > 9 && countAppearance < 12{
		return 1.00
	}
	if countAppearance >= 12{
		return 5.0
	}
	return 0.0
}



func CheckPlumPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 0.80
	}
	if countAppearance > 9 && countAppearance < 12{
		return 1.20
	}
	if countAppearance >= 12{
		return 8.0
	}
	return 0.0
}

func CheckApplePlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 1.00
	}
	if countAppearance > 9 && countAppearance < 12{
		return 1.50
	}
	if countAppearance >= 12{
		return 10.0
	}
	return 0.0
}

func CheckBlueCandyPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 1.50
	}
	if countAppearance > 9 && countAppearance < 12{
		return 2.00
	}
	if countAppearance >= 12{
		return 12.0
	}
	return 0.0
}

func CheckGreenCandyPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 2.00
	}
	if countAppearance > 9 && countAppearance < 12{
		return 5.00
	}
	if countAppearance >= 12{
		return 15.0
	}
	return 0.0
}


func CheckPurpleCandyPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 2.50
	}
	if countAppearance > 9 && countAppearance < 12{
		return 10.00
	}
	if countAppearance >= 12{
		return 25.0
	}
	return 0.0
}


func CheckRedCandyPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance > 7 && countAppearance < 10{
		return 10.0
	}
	if countAppearance > 9 && countAppearance < 12{
		return 25.0
	}
	if countAppearance >= 12{
		return 50.0
	}
	return 0.0
}

func CheckScatterPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	if countAppearance == 4{
		return 3.0
	}
	if countAppearance == 5{
		return 5.0
	}
	if countAppearance == 6{
		return 100.0
	}
	return 0.0
}


func CheckBomb2XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 2.0
}


func CheckBomb3XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 3.0
}

func CheckBomb5XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 5.0
}


func CheckBomb10XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 10.0
}

func CheckBomb25XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 25.0
}

func CheckBomb50XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 50.0
}

func CheckBomb100XPlayed(playingField [][]int, symbol int) float64{
	countAppearance := 0
	for i := 0; i < 5; i++{
		for j := 0; j < 6; j++{
			if playingField[i][j] == symbol{
				countAppearance += 1
			}
		}
	}
	return float64(countAppearance) * 100.0
}


func CalculatePaymentsNormalMode(playingField [][]int, bet float64, balance float64) float64{
	// Checking if symbols played
	payoutBananasX := CheckBananasPlayed(playingField, 1)
	payoutGrapesX := CheckGrapesPlayed(playingField, 2)
	payoutWatermelonX := CheckWatermelonPlayed(playingField, 3)
	payoutPlumX := CheckPlumPlayed(playingField, 4)
	payoutAppleX := CheckApplePlayed(playingField, 5)
	payoutBlueCandyX := CheckBlueCandyPlayed(playingField, 6)
	payoutGreenCandyX := CheckGreenCandyPlayed(playingField, 7)
	payoutPurpleCandyX := CheckPurpleCandyPlayed(playingField, 8)
	payoutRedCandyX := CheckRedCandyPlayed(playingField, 9)
	payoutScatterX := CheckScatterPlayed(playingField, 10)
	
    // Calculate total payout
    totalPayout := 0.0
    if payoutBananasX > 0.0 {
        totalPayout += bet * payoutBananasX
    }
    if payoutGrapesX > 0.0 {
        totalPayout += bet * payoutGrapesX
    }
    if payoutWatermelonX > 0.0 {
        totalPayout += bet * payoutWatermelonX
    }
    if payoutPlumX > 0.0 {
        totalPayout += bet * payoutPlumX
    }
    if payoutAppleX > 0.0 {
        totalPayout += bet * payoutAppleX
    }
    if payoutBlueCandyX > 0.0 {
        totalPayout += bet * payoutBlueCandyX
    }
    if payoutGreenCandyX > 0.0 {
        totalPayout += bet * payoutGreenCandyX
    }
    if payoutPurpleCandyX > 0.0 {
        totalPayout += bet * payoutPurpleCandyX
    }
    if payoutRedCandyX > 0.0 {
        totalPayout += bet * payoutRedCandyX
    }
    if payoutScatterX > 0.0 {
        totalPayout += bet * payoutScatterX
    }


	// New balance = Initial balance − Bet + Win
	balance = balance - bet + totalPayout

	return balance

}


func CalculatePaymentsBonusMode(playingField [][]int, bet float64, balance float64) float64{
	payoutBananasX := CheckBananasPlayed(playingField, 1)
	payoutGrapesX := CheckGrapesPlayed(playingField, 2)
	payoutWatermelonX := CheckWatermelonPlayed(playingField, 3)
	payoutPlumX := CheckPlumPlayed(playingField, 4)
	payoutAppleX := CheckApplePlayed(playingField, 5)
	payoutBlueCandyX := CheckBlueCandyPlayed(playingField, 6)
	payoutGreenCandyX := CheckGreenCandyPlayed(playingField, 7)
	payoutPurpleCandyX := CheckPurpleCandyPlayed(playingField, 8)
	payoutRedCandyX := CheckRedCandyPlayed(playingField, 9)
	payoutScatterX := CheckScatterPlayed(playingField, 10)
	payoutBomb2X := CheckBomb2XPlayed(playingField, 11)
	payoutBomb3X := CheckBomb3XPlayed(playingField, 12)
	payoutBomb5X := CheckBomb5XPlayed(playingField, 13)
	payoutBomb10X := CheckBomb10XPlayed(playingField, 14)
	payoutBomb25X := CheckBomb25XPlayed(playingField, 15)
	payoutBomb50X := CheckBomb50XPlayed(playingField, 16)
	payoutBomb100X := CheckBomb100XPlayed(playingField, 17)


	totalPayout := 0.0
    if payoutBananasX > 0.0 {
        totalPayout += bet * payoutBananasX
    }
    if payoutGrapesX > 0.0 {
        totalPayout += bet * payoutGrapesX
    }
    if payoutWatermelonX > 0.0 {
        totalPayout += bet * payoutWatermelonX
    }
    if payoutPlumX > 0.0 {
        totalPayout += bet * payoutPlumX
    }
    if payoutAppleX > 0.0 {
        totalPayout += bet * payoutAppleX
    }
    if payoutBlueCandyX > 0.0 {
        totalPayout += bet * payoutBlueCandyX
    }
    if payoutGreenCandyX > 0.0 {
        totalPayout += bet * payoutGreenCandyX
    }
    if payoutPurpleCandyX > 0.0 {
        totalPayout += bet * payoutPurpleCandyX
    }
    if payoutRedCandyX > 0.0 {
        totalPayout += bet * payoutRedCandyX
    }
    if payoutScatterX > 0.0 {
        totalPayout += bet * payoutScatterX
    }


	totalBomb := 0.0

	if payoutBomb2X > 0.0{
		totalBomb += payoutBomb2X
	}
	if payoutBomb3X > 0.0{
		totalBomb += payoutBomb3X
	}
	if payoutBomb5X > 0.0{
		totalBomb += payoutBomb5X
	}
	if payoutBomb10X > 0.0{
		totalBomb += payoutBomb10X
	}
	if payoutBomb25X > 0.0{
		totalBomb += payoutBomb25X
	}
	if payoutBomb50X > 0.0{
		totalBomb += payoutBomb50X
	}
	if payoutBomb100X > 0.0{
		totalBomb += payoutBomb100X
	}

	finalPayout := totalPayout * totalBomb

	// New balance = Initial balance + Total Win
	balance = balance + finalPayout
	return balance
} 

func WeightedRandomChoice(values []int, weights []float64) int {
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
    rand.Seed(time.Now().UnixNano())
    r := rand.Float64() * cumSum[len(cumSum)-1]

    // Находим элемент, соответствующий случайному числу
    for i, cs := range cumSum {
        if r < cs {
            return values[i]
        }
    }

    return values[len(values)-1] 
}

func GenerateRandomNumberNormalMode() int{
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    weights := []float64{10.0, 50.0, 100.0, 150.0, 200.0, 250.0, 300.0, 350.0, 450.0, 500.0}
	randomNumber := WeightedRandomChoice(values, weights)

}

func GenerateRandomNumberBonusMode(){
	rand.Seed(time.Now().UnixNano())

}


func GeneratePlayingField(){

}


func SweetBonanzaSpin(){

}