package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/gin-gonic/gin"
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