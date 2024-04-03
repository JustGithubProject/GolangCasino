package services

import (
	"errors"
	"strings"
	"math/rand"
	"time"
	"strconv"
    "net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
    "github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
    
)

func ShuffleWeights(arr []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}


func InitWeights(arr []int, length int){
	counter_weight := 10

	for i := 0; i < length; i++{
		arr[i] = counter_weight
		counter_weight += 100
	}
}


func InitNumbersArray(arr []int, length int){
	for i := 0; i < 37; i++{
		arr[i] = i
	}
}



func ValidateToken(c *gin.Context) (uint, error) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return 0, errors.New("Authorization header is missing")
    }

    tokenString := strings.Split(authHeader, " ")[1]
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("your_secret_key"), nil
    })
    if err != nil || !token.Valid {
        return 0, errors.New("Invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, errors.New("Invalid token claims")
    }
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("Invalid user ID in token")
    }
    userID := uint(userIDFloat)

    return userID, nil
}



// Функция для получения параметров игры из запроса
func GetGameParams(c *gin.Context) (string, int, float64, string, error) {
    guessSector := c.PostForm("guess_sector")
    guessNumber := c.PostForm("guess_number")
    bet := c.PostForm("bet")
    gameName := c.PostForm("gameName")

    guessNumberInt, err := strconv.Atoi(guessNumber)
    if err != nil {
        return "", 0, 0, "", errors.New("Invalid guess number")
    }

    betFloat, err := strconv.ParseFloat(bet, 64)
    if err != nil {
        return "", 0, 0, "", errors.New("Invalid bet")
    }

    return guessSector, guessNumberInt, betFloat, gameName, nil
}


func HandleGameRequest(c *gin.Context, fairPlay bool) {
    // Получаем JWT токен из заголовка запроса
    userID, err := ValidateToken(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }

    db := database.InitDB()
    user_repository := repositories.UserRepository{Db: db}

    user, err := user_repository.GetUserById(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    user_player := services.UserPlayer{}
    user_player.Balance = user.Balance

    guessSector, guessNumberInt, betFloat, gameName, err := GetGameParams(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game parameters"})
        return
    }

    if fairPlay {
        user_player.NormalPlay(guessSector, guessNumberInt, betFloat, gameName)
    } else {
        user_player.UnFairPlay(guessSector, guessNumberInt, betFloat, gameName)
    }
}