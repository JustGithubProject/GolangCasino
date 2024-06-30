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
	Bonbonnieres int // 11
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


func CheckBananasPlayed(playingField [][]int, symbol int) float32{
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

func CheckGrapesPlayed(playingField [][]int, symbol int) float32{
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


func CheckWatermelonPlayed(playingField [][]int, symbol int) float32{
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



func CheckPlumPlayed(playingField [][]int, symbol int) float32{
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

func CheckApplePlayed(playingField [][]int, symbol int) float32{
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

func CheckBlueCandyPlayed(playingField [][]int, symbol int) float32{
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

func CheckGreenCandyPlayed(playingField [][]int, symbol int) float32{
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


func CheckPurpleCandyPlayed(playingField [][]int, symbol int) float32{
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


func CheckRedCandyPlayed(playingField [][]int, symbol int) float32{
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

func CheckScatterPlayed(playingField [][]int, symbol int) float32{
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



func CalculatePayments(playingField [][]int){
	// Checking if the bananas played
	payoutX := CheckBananasPlayed(playingField, 1)
	// ...
}