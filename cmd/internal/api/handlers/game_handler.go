package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)


func CreateGameHandler(c *gin.Context){
	var game models.Game
	err_1 := c.BindJSON(&game)
	if err_1 != nil{
		c.JSON(400, gin.H{"error": err_1.Error()})
	}

	db := database.InitDB()

	gameRepository := repositories.GameRepository{Db: db}

	game_1 := models.Game{
		UserID: game.UserID,
		User: game.User,
		GameType: game.GameType,
		BetAmount: game.BetAmount,
		WinAmount: game.WinAmount,
		PlayedAt: game.PlayedAt,
	}

	err_2 := gameRepository.CreateGame(&game_1)
	if err_2 != nil{
		panic(err_2)
	}
}

func GetGameByIdHandler(c *gin.Context){
	gameIDStr := c.Param("id")
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
        return
	}
	db := database.InitDB()


	gameRepository := repositories.GameRepository{Db: db}
	game, err := gameRepository.GetGameById(uint(gameID))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
        return
	}

	c.JSON(http.StatusOK, game)

}


func UpdateGameHandler(c *gin.Context){
	gameIDStr := c.Param("id")
	gameID, err := strconv.Atoi(gameIDStr)
	
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
        return
	}

	db := database.InitDB()
	gameRepository := repositories.GameRepository{Db: db}
	modelGame, err := gameRepository.GetGameById(uint(gameID))
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
        return
    }

	// Обновляем данные игры
	err = gameRepository.UpdateGame(modelGame)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}