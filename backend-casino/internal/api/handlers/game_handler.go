package handlers

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
    "github.com/JustGithubProject/GolangCasino/backend-casino/internal/services/slots"
	"github.com/gin-gonic/gin"
)


func SpinRouletteHandler(c *gin.Context) {
    // Roulette handler
    services.HandleGameRequest(c, true)
}

func UnfairSpinRouletteHandler(c *gin.Context){
    // Unfair roulette handler
    services.HandleGameRequest(c, false)
}

func VeryBadSpinRouletteHandler(c *gin.Context){
    // Very bad spin roulette
    services.HandleVeryBadGameRequest(c)
}

func CreateGameHandler(c *gin.Context) {
    // Create Game Handler
    services.HandleCreateGame(c)
}


func GetGameByIdHandler(c *gin.Context) {
    // Handler to get game by id
    services.HandleGetGameByID(c)
}


func UpdateGameHandler(c *gin.Context) {
    // Handler to update game
    services.HandleUpdateGame(c)
}


func DeleteGameHandler(c *gin.Context) {
    // Handler to delete game
    services.HandleDeleteGame(c)
}



/*
    Slots
*/
func SweetBonanzaSlotHandler(c *gin.Context){
    // Sweetbonanza handler
    slots.SweetBonanzaHandle(c)
}

func DogHouseSlotHandler(c *gin.Context) {
    // Doghouse handler
    slots.DogHouseHandle(c)
}

func WolfGoldSlotHandler(c *gin.Context){
    // Wolfgold handler
    slots.WolfGoldHandle(c)
}

func BigBassBonanzaSlotHandler(c *gin.Context) {
    // Bigbassbonanza handler
    slots.BigBassBonanzaHandle(c)
}

func DiamondStrikeSlotHandler(c *gin.Context) {
    // Diamondstrike handler
    slots.DiamondStrikeHandle(c)
}