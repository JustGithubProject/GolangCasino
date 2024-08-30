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

func CheckBigBassBonanzaTenSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckBigBassBonanzaMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		if countSymbolsOnMiddleLine == 3 {
			return 0.50
		}

		if countSymbolsOnMiddleLine == 4 {
			return 2.50
		}

		if countSymbolsOnMiddleLine == 5 {
			return 10.00
		}
	}


	// Checking Upper Line
	countSymbolsOnUpperLine := CheckBigBassBonanzaUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		if countSymbolsOnUpperLine == 3 {
			return 0.50
		}

		if countSymbolsOnUpperLine == 4 {
			return 2.50
		}

		if countSymbolsOnUpperLine == 5 {
			return 10.00
		}
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckBigBassBonanzaLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		if countSymbolsOnLowerLine == 3 {
			return 0.50
		}

		if countSymbolsOnLowerLine == 4 {
			return 2.50
		}

		if countSymbolsOnLowerLine == 5 {
			return 10.00
		}
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckBigBassBonanzaFourthWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		if countSymbolsOnFourthLine == 3 {
			return 0.50
		}

		if countSymbolsOnFourthLine == 4 {
			return 2.50
		}

		if countSymbolsOnFourthLine == 5 {
			return 10.00
		}
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckBigBassBonanzaFifthWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		if countSymbolsOnFifthLine == 3 {
			return 0.50
		}

		if countSymbolsOnFifthLine == 4 {
			return 2.50
		}

		if countSymbolsOnFifthLine == 5 {
			return 10.00
		}
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckBigBassBonanzaSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		if countSymbolsOnSixthLine == 3 {
			return 0.50
		}

		if countSymbolsOnSixthLine == 4 {
			return 2.50
		}

		if countSymbolsOnSixthLine == 5 {
			return 10.00
		}
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckBigBassBonanzaSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		if countSymbolsOnSeventhLine == 3 {
			return 0.50
		}

		if countSymbolsOnSeventhLine == 4 {
			return 2.50
		}

		if countSymbolsOnSeventhLine == 5 {
			return 10.00
		}
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckBigBassBonanzaEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		if countSymbolsOnEighthLine == 3 {
			return 0.50
		}

		if countSymbolsOnEighthLine == 4 {
			return 2.50
		}

		if countSymbolsOnEighthLine == 5 {
			return 10.00
		}
	}

	return 0.0
}