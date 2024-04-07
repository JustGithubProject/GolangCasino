package services

import (
	"database/sql"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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


func InitNumbersArray(arr []int){
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

    user_player := UserPlayer{}
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


func HandleCreateGame(c *gin.Context){
    // Parse the JSON data from the request body into the game model
    var game models.Game
    err := c.BindJSON(&game)
    if err != nil {
        // If there's an error parsing JSON, return a 400 Bad Request response
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    gameRepository := repositories.GameRepository{Db: db}

    // Create a new game object with the parsed data
    game_1 := models.Game{
        UserID:    game.UserID,
        User:      game.User,
        GameType:  game.GameType,
        BetAmount: game.BetAmount,
        WinAmount: game.WinAmount,
        PlayedAt:  game.PlayedAt,
    }

    // Call the repository method to create the game in the database
    err = gameRepository.CreateGame(&game_1)
    if err != nil {
        panic(err)
    }
}


func HandleGetGameByID(c *gin.Context){
    // Extract the game ID from the request parameters
    gameIDStr := c.Param("id")
    gameID, err := strconv.Atoi(gameIDStr)
    if err != nil {
        // If the game ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    gameRepository := repositories.GameRepository{Db: db}

    // Call the repository method to get the game by its ID
    game, err := gameRepository.GetGameById(uint(gameID))
    if err != nil {
        // If there's an error getting the game, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
        return
    }

    // Return the game object as JSON
    c.JSON(http.StatusOK, game)
}


func HandleUpdateGame(c *gin.Context){
    // Extract the game ID from the request parameters
    gameIDStr := c.Param("id")
    gameID, err := strconv.Atoi(gameIDStr)
    if err != nil {
        // If the game ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    gameRepository := repositories.GameRepository{Db: db}

    // Call the repository method to get the game by its ID
    modelGame, err := gameRepository.GetGameById(uint(gameID))
    if err != nil {
        // If there's an error getting the game, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
        return
    }

    // Call the repository method to update the game
    err = gameRepository.UpdateGame(modelGame)
    if err != nil {
        // If there's an error updating the game, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
        return
    }

    // Return a success message
    c.JSON(http.StatusOK, gin.H{"message": "Game updated successfully"})
}


func HandleDeleteGame(c *gin.Context){
    // Extract the game ID from the request parameters
    gameIDStr := c.Param("id")
    gameID, err := strconv.Atoi(gameIDStr)
    if err != nil {
        // If the game ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    gameRepository := repositories.GameRepository{Db: db}

    // Call the repository method to get the game by its ID
    modelGame, err := gameRepository.GetGameById(uint(gameID))
    if err != nil {
        // If there's an error getting the game, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
        return
    }

    // Call the repository method to delete the game
    err = gameRepository.DeleteGame(modelGame)
    if err != nil {
        // If there's an error deleting the game, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
        return
    }

    // Return a success message
    c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}


func HandleUserRegister(c *gin.Context){
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := CreateUser(input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func HandleUserLogin(c *gin.Context){
    var userInput services.UserInput
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}
	user, err := userRepository.GetUserByEmail(userInput.Email)

	if err != nil {
        // Проверяем, что ошибка не связана с отсутствием пользователя
        if err == sql.ErrNoRows {
            // Если пользователь не найден, возвращаем ошибку 404
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        // Если возникла другая ошибка, возвращаем ошибку 400
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }
	if CheckPasswordHash(userInput.Password, user.Password){
		tokenString, err := CreateToken(userInput.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		// Устанавливаем заголовок Authorization
		c.Header("Authorization", "Bearer " + tokenString)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}