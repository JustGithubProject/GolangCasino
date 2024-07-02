package slots

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


func CalculatePayments(playingField [][]int, bet float64, balance float64) float64{
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
	

	// New balance = Initial balance − Bet + Win
	if payoutBananasX > 0.0{
		balance = balance - bet + (bet * payoutBananasX) 
	}
	if payoutGrapesX > 0.0{
		balance = balance - bet + (bet * payoutGrapesX)
	}
	if payoutWatermelonX > 0.0{
		balance = balance - bet + (bet * payoutWatermelonX)
	}
	if payoutPlumX > 0.0{
		balance = balance - bet + (bet * payoutPlumX)
	}
	if payoutAppleX > 0.0 {
		balance = balance - bet + (bet * payoutAppleX)
	}
	if payoutBlueCandyX > 0.0{
		balance = balance - bet + (bet * payoutBlueCandyX)
	}
	if payoutGreenCandyX > 0.0{
		balance = balance - bet + (bet * payoutGreenCandyX)
	}
	if payoutPurpleCandyX > 0.0{
		balance = balance - bet + (bet * payoutPurpleCandyX)
	}
	if payoutRedCandyX > 0.0{
		balance = balance - bet + (bet * payoutRedCandyX)
	}
	if payoutScatterX > 0.0{
		balance = balance - bet + (bet * payoutScatterX)
	}
	return balance

}


func BonusMode(playingField [][]int, bet float64){

} 