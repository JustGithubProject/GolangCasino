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

    // Getting a specific array from a matrix
    row1 := currentPlayingField[0]
    row2 := currentPlayingField[1]
    row3 := currentPlayingField[2]
    row4 := currentPlayingField[3]
    row5 := currentPlayingField[4]

    // Getting a specific string from an array
    sRow1 := ConversionArrayToString(row1)
    sRow2 := ConversionArrayToString(row2)
    sRow3 := ConversionArrayToString(row3)
    sRow4 := ConversionArrayToString(row4)
    sRow5 := ConversionArrayToString(row5)

    log.Println("Current playing field: ", currentPlayingField)

    // Update balance of user after spin
    user.Balance = currentBalance
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Game request handled successfully",
        "user": user,
        "balance": currentBalance,
        "sRow1": sRow1,
        "sRow2": sRow2,
        "sRow3": sRow3,
        "sRow4": sRow4,
        "sRow5": sRow5,
    })
}

