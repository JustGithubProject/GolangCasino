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