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