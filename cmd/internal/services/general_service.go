package services

import (
    "fmt"
	"database/sql"
	"errors"
	"html/template"
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

func convertStringToFloat64(rouletteElement string) (float64, error){
    convertedString, err := strconv.ParseFloat(rouletteElement, 64)
    if err != nil{
        fmt.Println("Failure to convert string to float64", err)
        return 0.0, err
    }
    return convertedString, nil

}


// Функция для получения параметров игры из запроса
func GetGameParams(c *gin.Context) (
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    float64,
    error) {
    guessEvenBet := c.PostForm("even")
    guessOddBet := c.PostForm("odd")
    guessRedBet := c.PostForm("red")
    guessBlackBet := c.PostForm("black")
    guessSectorBet := c.PostForm("sector")
    guessNumberBet := c.PostForm("number")
    guessOneToEighteenBet := c.PostForm("1To18")
    guessNineteenToThirtySixBet := c.PostForm("19To36")
    guessFirst2To1Bet := c.PostForm("First2To1")
    guessSecond2To1Bet := c.PostForm("Second2To1")
    guessThird2To1Bet := c.PostForm("Third2To1")

    guessEvenBet, err := convertStringToFloat64(guessEvenBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessOddBet, err := convertStringToFloat64(guessOddBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessRedBet, err := convertStringToFloat64(guessRedBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessBlackBet, err := convertStringToFloat64(guessBlackBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessSectorBet, err := convertStringToFloat64(guessSectorBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessNumberBet, err := convertStringToFloat64(guessNumberBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessOneToEighteenBet, err := convertStringToFloat64(guessOneToEighteenBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessNineteenToThirtySixBet, err := convertStringToFloat64(guessNineteenToThirtySixBetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessFirst2To1Bet, err := convertStringToFloat64(guessFirst2To1BetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessSecond2To1Bet, err := convertStringToFloat64(guessSecond2To1BetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }
    guessThird2To1Bet, err := convertStringToFloat64(guessThird2To1BetStr)
    if err != nil {
        return 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, err
    }

    return guessEvenBet, guessOddBet, guessRedBet, guessBlackBet, guessSectorBet, guessNumberBet, guessOneToEighteenBet, guessNineteenToThirtySixBet, guessFirst2To1Bet, guessSecond2To1Bet, guessThird2To1Bet
}

func InitializeUserRepository() (repositories.UserRepository, error){
    db := database.InitDB()
    user_repository := repositories.UserRepository{Db: db}
    return user_repository, nil
}

func InitializeGameRepository() (repositories.GameRepository, error){
    db := database.InitDB()
    game_repository := repositories.GameRepository{Db: db}
    return game_repository, nil
}


func HandleGameRequest(c *gin.Context, fairPlay bool) {
    // Получаем JWT токен из заголовка запроса
    userID, err := ValidateToken(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }

    user_repository, err := InitializeUserRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

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
    gameRepository, err := InitializeGameRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

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
    gameRepository, err := InitializeGameRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

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
    gameRepository, err := InitializeGameRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

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
    gameRepository, err := InitializeGameRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

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
    var userInput UserInput
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

    userRepository, err := InitializeUserRepository()
    if err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }
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

func HandleTemplate(c *gin.Context){
    // Loading the contents of the HTML file
    tmpl, err := template.ParseFiles("D:/Users/Kropi/Desktop/All directory/go/casino/cmd/templates/index.html")
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    
    // Send HTML page in response
    err = tmpl.Execute(c.Writer, nil)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
}