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



/*
	CHECK LEMON SYMBOL
*/
func CheckDiamondStrikeLemonSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.33, 0.66, 2.67)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.33, 0.66, 2.67)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.33, 0.66, 2.67)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.33, 0.66, 2.67)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.33, 0.66, 2.67)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.33, 0.66, 2.67)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.33, 0.66, 2.67)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.33, 0.66, 2.67)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.33, 0.66, 2.67)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.33, 0.66, 2.67)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.33, 0.66, 2.67)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.33, 0.66, 2.67)
	}

	return 0.0
}


/*
	CHECK PLUM SYMBOL
*/
func CheckDiamondStrikePlumSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.33, 0.66, 2.67)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.33, 0.66, 2.67)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.33, 0.66, 2.67)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.33, 0.66, 2.67)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.33, 0.66, 2.67)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.33, 0.66, 2.67)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.33, 0.66, 2.67)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.33, 0.66, 2.67)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.33, 0.66, 2.67)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.33, 0.66, 2.67)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.33, 0.66, 2.67)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.33, 0.66, 2.67)
	}

	return 0.0
}


/*
	CHECK WATERMELON SYMBOL
*/
func CheckDiamondStrikeWatermelonSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.33, 0.66, 2.67)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.33, 0.66, 2.67)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.33, 0.66, 2.67)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.33, 0.66, 2.67)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.33, 0.66, 2.67)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.33, 0.66, 2.67)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.33, 0.66, 2.67)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.33, 0.66, 2.67)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.33, 0.66, 2.67)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.33, 0.66, 2.67)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.33, 0.66, 2.67)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.33, 0.66, 2.67)
	}

	return 0.0
}


/*
	CHECK CHERRY SYMBOL
*/
func CheckDiamondStrikeCherrySymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.33, 0.66, 2.67)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.33, 0.66, 2.67)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.33, 0.66, 2.67)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.33, 0.66, 2.67)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.33, 0.66, 2.67)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.33, 0.66, 2.67)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.33, 0.66, 2.67)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.33, 0.66, 2.67)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.33, 0.66, 2.67)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.33, 0.66, 2.67)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.33, 0.66, 2.67)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.33, 0.66, 2.67)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.33, 0.66, 2.67)
	}

	return 0.0
}


/*
	CHECK BELL SYMBOL
*/
func CheckDiamondStrikeBellSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.66, 1.33, 6.67)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.66, 1.33, 6.67)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.66, 1.33, 6.67)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.66, 1.33, 6.67)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.66, 1.33, 6.67)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.66, 1.33, 6.67)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.66, 1.33, 6.67)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.66, 1.33, 6.67)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.66, 1.33, 6.67)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.66, 1.33, 6.67)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.66, 1.33, 6.67)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.66, 1.33, 6.67)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.66, 1.33, 6.67)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.66, 1.33, 6.67)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.66, 1.33, 6.67)
	}

	return 0.0
}


/*
	CHECK SEVEN SYMBOL
*/
func CheckDiamondStrikeSevenSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 0.66, 1.33, 13.33)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 0.66, 1.33, 13.33)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 0.66, 1.33, 13.33)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 0.66, 1.33, 13.33)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 0.66, 1.33, 13.33)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 0.66, 1.33, 13.33)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 0.66, 1.33, 13.33)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 0.66, 1.33, 13.33)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 0.66, 1.33, 13.33)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 0.66, 1.33, 13.33)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 0.66, 1.33, 13.33)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 0.66, 1.33, 13.33)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 0.66, 1.33, 13.33)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 0.66, 1.33, 13.33)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 0.66, 1.33, 13.33)
	}

	return 0.0
}


/*
	CHECK DIAMOND SYMBOL
*/
func CheckDiamondStrikeDiamondSymbolPlayed(playingField [][]int, symbol int) float64 {
	// Checking Middle Line
	countSymbolsOnMiddleLine := CheckDiamondStrikeMiddleWinLine(playingField, symbol)
	if countSymbolsOnMiddleLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnMiddleLine, 1.33, 4.00, 20.00)
	}

	// Checking Upper Line
	countSymbolsOnUpperLine := CheckDiamondStrikeUpperWinLine(playingField, symbol)
	if countSymbolsOnUpperLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnUpperLine, 1.33, 4.00, 20.00)
	}

	// Checking Lower Line
	countSymbolsOnLowerLine := CheckDiamondStrikeLowerWinLine(playingField, symbol)
	if countSymbolsOnLowerLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnLowerLine, 1.33, 4.00, 20.00)
	}

	// Checking Fourth Line
	countSymbolsOnFourthLine := CheckDiamondStrikeVWinLine(playingField, symbol)
	if countSymbolsOnFourthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourthLine, 1.33, 4.00, 20.00)
	}

	// Checking Fifth Line
	countSymbolsOnFifthLine := CheckDiamondStrikeReverseVWinLine(playingField, symbol)
	if countSymbolsOnFifthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifthLine, 1.33, 4.00, 20.00)
	}

	// Checking Sixth Line
	countSymbolsOnSixthLine := CheckDiamondStrikeSixthWinLine(playingField, symbol)
	if countSymbolsOnSixthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSixthLine, 1.33, 4.00, 20.00)
	}

	// Checking Seventh Line
	countSymbolsOnSeventhLine := CheckDiamondStrikeSeventhWinLine(playingField, symbol)
	if countSymbolsOnSeventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnSeventhLine, 1.33, 4.00, 20.00)
	}

	// Checking Eighth Line
	countSymbolsOnEighthLine := CheckDiamondStrikeEighthWinLine(playingField, symbol)
	if countSymbolsOnEighthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEighthLine, 1.33, 4.00, 20.00)
	}

	// Checking Ninth Line
	countSymbolsOnNinthLine := CheckDiamondStrikeNinthWinLine(playingField, symbol)
	if countSymbolsOnNinthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnNinthLine, 1.33, 4.00, 20.00)
	}

	// Checking Tenth Line
	countSymbolsOnTenthLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnTenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTenthLine, 1.33, 4.00, 20.00)
	}

	// Checking 11-th Line
	countSymbolsOnEleventhLine := CheckDiamondStrikeTenthWinLine(playingField, symbol)
	if countSymbolsOnEleventhLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnEleventhLine, 1.33, 4.00, 20.00)
	}

	// Checking 12-th Line
	countSymbolsOnTwelfthLine := CheckDiamondStriketTwelfthWinLine(playingField, symbol)
	if countSymbolsOnTwelfthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnTwelfthLine, 1.33, 4.00, 20.00)
	}

	// Checking 13-th Line
	countSymbolsOnThirteenthLine := CheckDiamondStriketThirteenthWinLine(playingField, symbol)
	if countSymbolsOnThirteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnThirteenthLine, 1.33, 4.00, 20.00)
	}

	// Checking 14-th Line
	countSymbolsOnFourteenthLine := CheckDiamondStriketFourteenthWinLine(playingField, symbol)
	if countSymbolsOnFourteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFourteenthLine, 1.33, 4.00, 20.00)
	}

	// Checking 15-th Line
	countSymbolsOnFifteenthLine := CheckDiamondStriketFifteenthWinLine(playingField, symbol)
	if countSymbolsOnFifteenthLine > 2 {
		return GreaterThanThreeSymbol(countSymbolsOnFifteenthLine, 1.33, 4.00, 20.00)
	}

	return 0.0
}

// type DiamondStrikeSymbols struct {
// 	Lemon int // symbol=Lemon order=1
// 	Plum int // symbol=Plum order=2
// 	Watermelon int // symbol=Watermelon order=3
// 	Cherry int // symbol=Cherry order=4
// 	Bell int // symbol=Bell order=5
// 	Seven int // symbol=Seven order=6
// 	Diamond int // symbol=Diamond order=7
// }


func CalculateDiamondStrikePaymentsNormalMode(playingField [][]int, bet float64, balance float64) float64 {
	// Checking the payout for each set of symbols
	payoutLemonX := CheckDiamondStrikeLemonSymbolPlayed(playingField, 1)
	payoutPlumX := CheckDiamondStrikePlumSymbolPlayed(playingField, 2)
	payoutWatermelonX := CheckDiamondStrikeWatermelonSymbolPlayed(playingField, 3)
	payoutCherryX := CheckDiamondStrikeCherrySymbolPlayed(playingField, 4)
	payoutBellX := CheckDiamondStrikeBellSymbolPlayed(playingField, 5)
	payoutSevenX := CheckDiamondStrikeSevenSymbolPlayed(playingField, 6)
	payoutDiamondX := CheckDiamondStrikeDiamondSymbolPlayed(playingField, 7)

	// Calculating total payout
	totalPayout := 0.0

	if payoutLemonX > 0.0 {
		totalPayout += bet * payoutLemonX
	}

	if payoutPlumX > 0.0 {
		totalPayout += bet * payoutPlumX
	}

	if payoutWatermelonX > 0.0 {
		totalPayout += bet * payoutWatermelonX
	}

	if payoutCherryX > 0.0 {
		totalPayout += bet * payoutCherryX
	}

	if payoutBellX > 0.0 {
		totalPayout += bet * payoutBellX
	}

	if payoutSevenX > 0.0 {
		totalPayout += bet * payoutSevenX
	}

	if payoutDiamondX > 0.0 {
		totalPayout += bet * payoutDiamondX
	}

	// New balance = Initial balance − Bet + Win
	balance = balance - bet + totalPayout
	return balance
}

func WeightedDiamondStrikeRandomChoice(values []int, weights []float64) int {
	if len(values) != len(weights) {
		panic("Values and weights must be the same length")
	}

	// Weight to invert
	invWeights := make([]float64, len(weights))
	for i, w := range weights {
		invWeights[i] = 1.0 / w
	}

	// Make a list of cumulative sums
	cumSum := make([]float64, len(invWeights))
	cumSum[0] = invWeights[0]
	for i := 1; i < len(invWeights); i++ {
		cumSum[i] = cumSum[i-1] + invWeights[i]
	}

	// Generate a random number
	r := rand.Float64() * cumSum[len(cumSum)-1]

	// Находим элемент, соответствующий случайному числу
	for i, cs := range cumSum {
		if r < cs {
			return values[i]
		}
	}

	return values[len(values)-1]
}

func GenerateDiamondStrikeRandomNumberNormalMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	weights := []float64{
		5.0, 5.0, 10.0, 15.0, 20.0, 25.0, 50.0,
	}
	randomNumber := WeightedDiamondStrikeRandomChoice(values, weights)
	return randomNumber
}

func GenerateDiamondStrikeRandomNumberBonusMode() int {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	weights := []float64{
		5.0, 5.0, 10.0, 15.0, 20.0, 25.0, 50.0,
	}
	randomNumber := WeightedDiamondStrikeRandomChoice(values, weights)
	return randomNumber
}

func init() {
	rand.Seed(time.Now().UnixNano())
}


func GenerateDiamondStrikePlayingFieldNormalMode() [][]int {
	cols, rows := 3, 5
	playingField := CreateDiamondStrikePlayingField()
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			playingField[i][j] = GenerateDiamondStrikeRandomNumberNormalMode()
		}
	}
	return playingField
}

func GenerateDiamondStrikePlayingFieldBonusMode() [][]int {
	playingField := CreateDiamondStrikePlayingField()
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			playingField[i][j] = GenerateDiamondStrikeRandomNumberBonusMode()
		}
	}
	return playingField
}

func DiamondStrikeSpin(isBonusMode bool, bet float64, balance float64) ([][]int, float64) {
	if isBonusMode {
		currentPlayingField := GenerateDiamondStrikePlayingFieldBonusMode()
		currentBalance := CalculateDiamondStrikePaymentsNormalMode(currentPlayingField, bet, balance) // Isn't correct.
		return currentPlayingField, currentBalance
	}
	currentPlayingField := GenerateDiamondStrikePlayingFieldNormalMode()
	currentBalance := CalculateDiamondStrikePaymentsNormalMode(currentPlayingField, bet, balance)
	return currentPlayingField, currentBalance
}