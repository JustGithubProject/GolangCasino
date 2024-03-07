package repositories

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"gorm.io/gorm"
)


type GameRepository struct {
	db *gorm.DB
}


// Method to create a game
func (gm *GameRepository) CreateGame(game *models.Game) error {
	return gm.db.Create(game).Error
}


// Method to get user by id
func (gm *GameRepository) GetGameById(id uint) (*models.Game, error){
	var game models.Game
	err := gm.db.First(&game, id).Error
	if err != nil{
		return nil, err
	}
	return &game, nil
}

// Method to update game
func (gm *GameRepository) UpdateGame(game *models.Game) error {
    return gm.db.Save(game).Error
}


// Method to delete game
func (gm *GameRepository) DeleteGame(game *models.Game) error {
    return gm.db.Delete(game).Error
}

