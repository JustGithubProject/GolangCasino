package slots

import (
	"math/rand"
	"time"
)

func CreateWolfGoldPlayingField() [][]int {
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

type WolfGoldSymbols struct {
	// Low
	jack int // symbol=J order=1
	queen int // symbol=Q order=2
	king int // symbol=K order=3
	ace int // symbol=A order=4

	// Middle
	cougar int // symbol=Cougar order=5
	horse int // symbol=Horse order=6

	// High
	eagle int // symbol=Eagle order=7
	bison int // symbol=Bison order=8
	wolf int // symbol=Wolf order=9

	// Scatter
	scatter int // symbol=Scatter order=10
}