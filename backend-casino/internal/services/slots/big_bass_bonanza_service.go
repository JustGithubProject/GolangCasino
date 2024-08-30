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

func CheckBigBassBonanzaUppperWinLine(playingField [][]int, symbol int) int {
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


func CheckBigBassBonanzaTenSymbolPlayed(playingField [][]int, symbol int) float64 {
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