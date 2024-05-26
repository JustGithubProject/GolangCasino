package services


// GameError represents a game-related error
type GameError struct {
	Message string
}

func (e *GameError) Error() string {
    return e.Message
}



