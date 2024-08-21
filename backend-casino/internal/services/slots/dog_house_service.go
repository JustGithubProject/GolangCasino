package slots

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

	// Bonus
	paws int // symbol=Paws order=12
}



/*
	BLOCK: Checking multipliers
*/

func CountMatrixElement(playingField[][] int, symbol int) int {
	countAppearance := 0

	for i := 0; i < 3; i++{
		for j := 0; j < 5; j++{
			if playingField[i][j] == symbol {
				countAppearance += 1
			}
		}
	}
	return countAppearance
}


func CheckTenSymbolPlayed(playingField[][]int, symbol int) float64{
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


func CheckJackSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

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
	countAppearance := CountMatrixElement(playingField, symbol)

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
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckAceSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckBoneSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckCollarSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckDachshundSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckPugSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckSpitzSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckBoxerSymbolPlayed(playingField[][]int, symbol int) float64 {
	countAppearance := CountMatrixElement(playingField, symbol)

	if countAppearance == 3{
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


func CheckPawsSymbolPlayed(playingField[][]int, symbol int) float64 {
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


func CalculateDogHousePaymentsNormalMode(playingField[][]int, bet float64, balance float64) float64 {
	payoutTenX := CheckTenSymbolPlayed(playingField, 1)
	payoutJackX := CheckJackSymbolPlayed(playingField, 2)
	payoutQueenX := CheckQueenSymbolPlayed(playingField, 3)
	payoutKingX := CheckKingSymbolPlayed(playingField, 4)
	payoutAceX := CheckAceSymbolPlayed(playingField, 5)
	payoutBoneX := CheckBoneSymbolPlayed(playingField, 6)
	payoutCollarX := CheckCollarSymbolPlayed(playingField, 7)
	payoutDachshundX := CheckDachshundSymbolPlayed(playingField, 8)
	payoutPugX := CheckPugSymbolPlayed(playingField, 9)
	payoutSpitz := CheckSpitzSymbolPlayed(playingField, 10)
	// TODO: ...



}