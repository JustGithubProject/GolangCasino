package slots

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
	"github.com/gin-gonic/gin"
)


func SweetBonanzaHandle(c *gin.Context) {
	username, err := services.ValidateToken(c)
    if err != nil {
        fmt.Println("С токеном проблемы?")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    user_repository, err := services.InitializeUserRepository()
    if err != nil{
        fmt.Println("С репозиторием ?")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

    user, err := user_repository.GetUserByUsername(username)
    fmt.Printf("User=%v+\n", user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }
	spinBetStr := c.Query("spinBet")
	convertedSpinBet, _ := strconv.ParseFloat(spinBetStr, 64)

	// we do not take first var because it's not bonus mode
	_, currentBalance := SweetBonanzaSpin(false, convertedSpinBet, user.Balance)

	user.Balance = currentBalance
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Game request handled successfully", "user": user})
}
