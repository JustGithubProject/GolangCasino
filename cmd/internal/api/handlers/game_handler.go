package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/services"
	"github.com/gin-gonic/gin"
)


func SpinRouletteHandler(c *gin.Context) {
    services.HandleGameRequest(c, true)
}


func UnfairSpinRouletteHandler(c *gin.Context){
    services.HandleGameRequest(c, false)
}


func CreateGameHandler(c *gin.Context) {
    services.HandleCreateGame(c)
}


func GetGameByIdHandler(c *gin.Context) {
    services.HandleGetGameByID(c)
}


func UpdateGameHandler(c *gin.Context) {
    services.HandleUpdateGame(c)
}


func DeleteGameHandler(c *gin.Context) {
    services.HandleDeleteGame(c)
}
