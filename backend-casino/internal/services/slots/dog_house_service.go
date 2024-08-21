package slots

// func CreateDogHousePlayingField() [][]int{
// 	rows, cols := 5, 6
// 	playingField := make([][]int, rows)
//     for i := range playingField {
//         playingField[i] = make([]int, cols)
//         for j := range playingField[i] {
//             playingField[i][j] = 0
//         }
//     }
// 	return playingField
// }

func CreateDogHousePlayingField() [][]int {
	rows, cols := 3, 5
	playingField := make([][]int, rows)
	for i := range playingField {
		playingField[i] = make([]int, cols)
		for j := range playingField[i] {
			playingField[i][j] = 0
		}
	}
	return playingField
}

type DogHouseSymbols struct {
	// Low value symbols
	ten int  // symbol=10 order=1
	jack int  // symbol=J order=2
	queen int  // symbol=Q order=3
	king int  // symbol=K order=4
	ace int  // symbol=A order=5

	// Medium value symbols
	bone int // symbol=Bone order=6
	collar int // symbol=Collar order=7

	// High Value symbols (Dogs)
	dachshund int // symbol=Dachshund order=8
	pug int // symbol=Pug order=9
	spitz int // symbol=Spitz order=10
	boxer int // symbol=Boxer order=11
}



/*
	BLOCK: Checking multipliers
*/
func CheckTenSymbolPlayed(playingField[][]int, symbol int) float64{
	countAppearance := 0

	for i := 0; i < 3; i++{
		for j := 0; j < 5; j++{
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}

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


func CheckJackSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := 0

	for i := 0; i < 3; i++{
		for j := 0; j < 5; j++ {
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}

	if countAppearance == 3{
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


func CheckQueenSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := 0

	for i := 0; i < 3; i++{
		for j := 0; j < 5; j++ {
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}

	if countAppearance == 3{
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


func CheckKingSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := 0

	for i := 0; i < 3; i++{
		for j := 0; j < 5; j++ {
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}

	if countAppearance == 3{
		return 0.25
	}

	if countAppearance == 4 {
		return 0.50
	}

	if countAppearance == 5 {
		// TODO: ...
		return 0.0
	}

	return 0.0
}