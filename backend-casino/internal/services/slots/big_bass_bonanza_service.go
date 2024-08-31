package slots

// import (
// 	"math/rand"
// 	"time"
// )

func CreateBigBassBonanzaPlayingField() [][]int {
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

type BigBassBonanzaSymbols struct {
	// Low value symbols
	ten   int // symbol=10 order=1
	jack  int // symbol=J order=2
	queen int // symbol=Q order=3
	king  int // symbol=K order=4
	ace   int // symbol=A order=5

	// Medium value symbols
	fish   int // symbol=Fish order=6
	boxes int // symbol=Boxes order=7

	// High Value symbols
	dragonfly int // symbol=Dragonfly order=8
	rod       int // symbol=Rod order=9
	bobber     int // symbol=Spitz order=10

	// Bonus
	fishScatter int // symbol=Paws order=11
}

func CheckBigBassBonanzaMiddleWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}
	return counter
}

func CheckBigBassBonanzaUpperWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}
	return counter
}

func CheckBigBassBonanzaLowerWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}
	return counter
}

func CheckBigBassBonanzaFourthWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 1; i < 4; i++ {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}

	if playingField[1][0] == symbol {
		counter += 1
	}

	if playingField[1][4] == symbol {
		counter += 1
	}

	return counter
}


func CheckBigBassBonanzaFifthWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 1; i < 4; i++ {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}

	if playingField[1][0] == symbol {
		counter += 1
	}
	
	if  playingField[1][4] == symbol {
		counter += 1
	}
	return counter
}


func CheckBigBassBonanzaSixthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][2] == symbol {
		counter += 1
	}

	if playingField[1][1] == symbol {
		counter += 1
	}

	if playingField[1][3] == symbol {
		counter += 1
	}

	if playingField[2][0] == symbol {
		counter += 1
	}

	if playingField[2][4] == symbol {
		counter += 1
	}

	return counter
}

func CheckBigBassBonanzaSeventhWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][0] == symbol {
		counter += 1
	}

	if playingField[0][4] == symbol {
		counter += 1
	}

	if playingField[1][1] == symbol {
		counter += 1
	}

	if playingField[1][3] == symbol {
		counter += 1
	}
	
	if playingField[2][2] == symbol {
		counter += 1
	}

	return counter
}

func CheckBigBassBonanzaEighthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][3] == symbol {
		counter += 1
	}

	if playingField[0][4] == symbol {
		counter += 1
	}

	if playingField[1][2] == symbol {
		counter += 1
	}

	if playingField[2][0] == symbol {
		counter += 1
	}

	if playingField[2][1] == symbol {
		counter += 1
	}

	return counter
}

func CheckBigBassBonanzaNinthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][0] == symbol {
		counter += 1
	}

	if playingField[0][1] == symbol {
		counter += 1
	}

	if playingField[1][2] == symbol {
		counter += 1
	}

	if playingField[2][3] == symbol {
		counter += 1
	}

	if playingField[2][4] == symbol {
		counter += 1
	}

	return counter
}

func CheckBigBassBonanzaTenthWinLine(playingField [][]int, symbol int) int {
	counter := 0
	
	for i := 1; i < 4; i++ {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	if playingField[0][4] == symbol {
		counter += 1
	}

	if playingField[2][0] == symbol {
		counter += 1
	}

	return counter
}


func GreaterThanThreeSymbol(countLine int, factorFor3 float64, factorFor4 float64, factorFor5 float64) float64 {
	if countLine == 3 {
		return factorFor3
	}

	if countLine == 4 {
		return factorFor4
	}

	if countLine == 5 {
		return factorFor5
	}

	return 0.0
}

/*
	CHECK TEN SYMBOL
*/
func CheckBigBassBonanzaTenSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.50, 2.50, 10.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.50, 2.50, 10.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.50, 2.50, 10.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.50, 2.50, 10.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.50, 2.50, 10.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.50, 2.50, 10.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.50, 2.50, 10.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.50, 2.50, 10.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.50, 2.50, 10.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.50, 2.50, 10.00)
	}
	return 0.0
}

/*
	CHECK JACK SYMBOL
*/
func CheckBigBassBonanzaJackSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.50, 2.50, 10.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.50, 2.50, 10.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.50, 2.50, 10.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.50, 2.50, 10.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.50, 2.50, 10.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.50, 2.50, 10.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.50, 2.50, 10.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.50, 2.50, 10.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.50, 2.50, 10.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.50, 2.50, 10.00)
	}
	return 0.0
}


/*
	CHECK QUEEN SYMBOL
*/
func CheckBigBassBonanzaQueenSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.50, 2.50, 10.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.50, 2.50, 10.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.50, 2.50, 10.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.50, 2.50, 10.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.50, 2.50, 10.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.50, 2.50, 10.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.50, 2.50, 10.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.50, 2.50, 10.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.50, 2.50, 10.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.50, 2.50, 10.00)
	}
	return 0.0
}


/*
	CHECK KING SYMBOL
*/
func CheckBigBassBonanzaKingSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.50, 2.50, 10.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.50, 2.50, 10.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.50, 2.50, 10.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.50, 2.50, 10.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.50, 2.50, 10.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.50, 2.50, 10.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.50, 2.50, 10.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.50, 2.50, 10.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.50, 2.50, 10.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.50, 2.50, 10.00)
	}
	return 0.0
}


/*
	CHECK ACE SYMBOL
*/
func CheckBigBassBonanzaAceSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.50, 2.50, 10.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.50, 2.50, 10.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.50, 2.50, 10.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.50, 2.50, 10.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.50, 2.50, 10.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.50, 2.50, 10.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.50, 2.50, 10.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.50, 2.50, 10.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.50, 2.50, 10.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.50, 2.50, 10.00)
	}
	return 0.0
}


/*
	CHECK FISH SYMBOL
*/
func CheckBigBassBonanzaFishSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 1.00, 5.00, 20.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 1.00, 5.00, 20.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 1.00, 5.00, 20.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 1.00, 5.00, 20.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 1.00, 5.00, 20.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 1.00, 5.00, 20.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 1.00, 5.00, 20.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 1.00, 5.00, 20.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 1.00, 5.00, 20.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 1.00, 5.00, 20.00)
	}
	return 0.0
}


/*
	CHECK BOXES SYMBOL
*/
func CheckBigBassBonanzaBoxesSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 2.00, 10.00, 50.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 2.00, 10.00, 50.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 2.00, 10.00, 50.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 2.00, 10.00, 50.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 2.00, 10.00, 50.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 2.00, 10.00, 50.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 2.00, 10.00, 50.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 2.00, 10.00, 50.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 2.00, 10.00, 50.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 2.00, 10.00, 50.00)
	}
	return 0.0
}


/*
	CHECK DRAGONFLY SYMBOL
*/
func CheckBigBassBonanzaDragonflySymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 2.00, 10.00, 50.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 2.00, 10.00, 50.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 2.00, 10.00, 50.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 2.00, 10.00, 50.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 2.00, 10.00, 50.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 2.00, 10.00, 50.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 2.00, 10.00, 50.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 2.00, 10.00, 50.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 2.00, 10.00, 50.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 2.00, 10.00, 50.00)
	}
	return 0.0
}


/*
	CHECK ROD SYMBOL
*/
func CheckBigBassBonanzaRodSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 3.00, 15.00, 100.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 3.00, 15.00, 100.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 3.00, 15.00, 100.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 3.00, 15.00, 100.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 3.00, 15.00, 100.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 3.00, 15.00, 100.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 3.00, 15.00, 100.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 3.00, 15.00, 100.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 3.00, 15.00, 100.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 3.00, 15.00, 100.00)
	}
	return 0.0
}


// CHECK BOBBER SYMBOL
func CheckBigBassBonanzaBobberSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine == 2 {
		return 0.50
	}
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 5.00, 20.00, 200.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine == 2 {
		return 0.50
	}
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 5.00, 20.00, 200.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine == 2 {
		return 0.50
	}
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 5.00, 20.00, 200.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine == 2 {
		return 0.50
	}
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 5.00, 20.00, 200.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine == 2 {
		return 0.50
	}
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 5.00, 20.00, 200.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine == 2 {
		return 0.50
	}
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 5.00, 20.00, 200.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine == 2 {
		return 0.50
	}
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 5.00, 20.00, 200.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine == 2 {
		return 0.50
	}
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 5.00, 20.00, 200.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckBigBassBonanzaNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine == 2 {
		return 0.50
	}
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 5.00, 20.00, 200.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckBigBassBonanzaTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine == 2 {
		return 0.50
	}
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 5.00, 20.00, 200.00)
	}
	return 0.0
}