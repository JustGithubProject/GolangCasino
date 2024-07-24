package slots

import (
	"net/http"
	"strconv"
    "log"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
	"github.com/gin-gonic/gin"
)


func SweetBonanzaHandle(c *gin.Context) {
    // Validation of token
    username, err := services.ValidateToken(c)
    if err != nil {
        log.Println("Issues with the token!")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    user_repository, err := services.InitializeUserRepository()
    if err != nil{
        log.Println("Failed to init repository")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

    // Getting user by username derived from token
    user, err := user_repository.GetUserByUsername(username)
    log.Printf("User=%v+\n", user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Take spinBet and convert to float
    spinBetStr := c.Query("spinBet")
    convertedSpinBet, _ := strconv.ParseFloat(spinBetStr, 64)

    // Do not allow the user to make a spin if they do not have enough money
    if user.Balance < convertedSpinBet{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have enough money"})
        return
    }

    currentPlayingField, currentBalance := SweetBonanzaSpin(false, convertedSpinBet, user.Balance)
    log.Println("Current playing field: ", currentPlayingField)

    // Update balance of user after spin
    user.Balance = currentBalance
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Game request handled successfully", "user": user, "playingField": currentPlayingField, "balance": currentBalance})
}

