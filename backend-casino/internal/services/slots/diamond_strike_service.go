package slots

import (
	"math/rand"
	"time"
)


func CreateDiamondStrikePlayingField() [][]int {
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


type DiamondStrikeSymbols struct {
	Lemon int // symbol=Lemon order=1
	Plum int // symbol=Plum order=2
	Watermelon int // symbol=Watermelon order=3
	Cherry int // symbol=Cherry order=4
	Bell int // symbol=Bell order=5
	Seven int // symbol=Seven order=6
	Diamond int // symbol=Diamond order=7
}


func CheckDiamondStrikeMiddleWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}
	return counter
}

func CheckDiamondStrikeUpperWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}
	return counter
}

func CheckDiamondStrikeLowerWinLine(playingField [][]int, symbol int) int {
	counter := 0
	for i := 0; i < 5; i++ {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}
	return counter
}


func CheckDiamondStrikeVWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first diagonal
	for i := 0; i < 3; i++ { // 0, 1, 2
		if playingField[i][i] == symbol {
			counter += 1
		}
	}

	// Checking second diagonal
	if playingField[0][4] == symbol {
		counter += 1
	}

	if playingField[1][3] == symbol {
		counter += 1
	}

	return counter
}


func CheckDiamondStrikeReverseVWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first diagonal
	j := 2
	for i := 0; i < 3; i++ {
		if playingField[i][j] == symbol {
			counter += 1
		}
		j -= 1
	}

	// Checking second diagonal
	if playingField[1][3] == symbol {
		counter += 1
	}

	if playingField[2][4] == symbol {
		counter += 1
	}

	return counter
}

func CheckDiamondStrikeSixthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking upper line except first and last cells
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

func CheckDiamondStrikeSeventhWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking lower line except first and last cells
	for i := 1; i < 4; i++{
		if playingField[2][i] == symbol {
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


func CheckDiamondStrikeEighthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first two cells of upper line
	for i := 0; i < 2; i++ {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}

	// Checking middle cell
	if playingField[1][2] == symbol {
		counter += 1
	}

	// Checking last two cells of lower line
	for i := 3; i < 5; i++ {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}

	return counter
}

func CheckDiamondStrikeNinthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first two cells of lower line
	for i := 0; i < 2; i++ {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}

	// Checking middle cell
	if playingField[1][2] == symbol {
		counter += 1
	}

	// Checking last two cells of upper line
	for i := 3; i < 5; i++ {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}

	return counter
}

func CheckDiamondStrikeTenthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][3] == symbol {
		counter += 1
	}

	// Checking the middle cell with a paired index
	for i := 0; i < 5; i += 2 {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	if playingField[2][1] == symbol {
		counter += 1
	}
	return counter
}

func CheckDiamondStrikeEleventhWinLine(playingField [][]int, symbol int) int {
	counter := 0

	if playingField[0][1] == symbol {
		counter += 1
	}

	// Checking the middle cell with a paired index
	for i := 0; i < 5; i += 2 {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	if playingField[2][3] == symbol {
		counter += 1
	}

	return counter
} 


func CheckDiamondStriketTwelfthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first and last cells of upper line
	for i := 0; i < 5; i += 4 {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}

	// Checking middle line except first and last cells
	for i := 1; i < 4; i++ {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	return counter
}


func CheckDiamondStriketThirteenthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking first and last cells of lower line
	for i := 0; i < 5; i += 4 {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}

	// Checking middle line except first and last cells
	for i := 1; i < 4; i++ {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	return counter
}


func CheckDiamondStriketFourteenthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking paired cells of upper line
	for i := 0; i < 5; i += 2 {
		if playingField[0][i] == symbol {
			counter += 1
		}
	}

	// Checking two cells of middle line 1-index and 3-index 
	for i := 1; i < 4; i += 2 {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	return counter
}


func CheckDiamondStriketFifteenthWinLine(playingField [][]int, symbol int) int {
	counter := 0

	// Checking two cells of middle line 1-index and 3-index 
	for i := 1; i < 4; i += 2 {
		if playingField[1][i] == symbol {
			counter += 1
		}
	}

	// Checking paired cells of lower line
	for i := 0; i < 5; i += 2 {
		if playingField[2][i] == symbol {
			counter += 1
		}
	}

	return counter
}