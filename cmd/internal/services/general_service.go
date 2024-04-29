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


type GameParams struct {
    GuessEvenBet             float64
    GuessOddBet              float64
    GuessRedBet              float64
    GuessBlackBet            float64
    GuessSector1st12Bet      float64
    GuessSector2nd12Bet      float64
    GuessSector3rd12Bet      float64
    GuessNumberBet           float64
    GuessOneToEighteenBet    float64
    GuessNineteenToThirtySix float64
    GuessFirst2To1Bet        float64
    GuessSecond2To1Bet       float64
    GuessThird2To1Bet        float64
    Err                      error
}


// Функция для получения параметров игры из запроса
func GetGameParams(c *gin.Context) GameParams {
    var (
        guessEvenBet             float64
        guessOddBet              float64
        guessRedBet              float64
        guessBlackBet            float64
        guessSector1st12Bet      float64
        guessSector2nd12Bet      float64
        guessSector3rd12Bet      float64
        guessNumberBet           float64
        guessOneToEighteenBet    float64
        guessNineteenToThirtySixBet float64
        guessFirst2To1Bet        float64
        guessSecond2To1Bet       float64
        guessThird2To1Bet        float64
        err                      error
    )
    guessEvenBetStr := c.PostForm("even")
    guessOddBetStr := c.PostForm("odd")
    guessRedBetStr := c.PostForm("red")
    guessBlackBetStr := c.PostForm("black")
    guessSector1st12BetStr := c.PostForm("1 st 12")
    guessSector2nd12BetStr := c.PostForm("2 nd 12")
    guessSector3rd12BetStr := c.PostForm("3 rd 12")
    guessNumberBetStr := c.PostForm("number")
    guessOneToEighteenBetStr := c.PostForm("1To18")
    guessNineteenToThirtySixBetStr := c.PostForm("19To36")
    guessFirst2To1BetStr := c.PostForm("First2To1")
    guessSecond2To1BetStr := c.PostForm("Second2To1")
    guessThird2To1BetStr := c.PostForm("Third2To1")

    // Функция для преобразования строки в float64
    convertStringToFloat64 := func(str string) (float64, error) {
        if str == "" {
            return 0, nil // Возвращаем 0, если строка пустая
        }
        return strconv.ParseFloat(str, 64)
    }

    // Преобразование и обработка значений
    guessEvenBet, err = convertStringToFloat64(guessEvenBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessOddBet, err = convertStringToFloat64(guessOddBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessRedBet, err = convertStringToFloat64(guessRedBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessBlackBet, err = convertStringToFloat64(guessBlackBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessSector1st12Bet, err = convertStringToFloat64(guessSector1st12BetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessSector2nd12Bet, err = convertStringToFloat64(guessSector2nd12BetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessSector3rd12Bet, err = convertStringToFloat64(guessSector3rd12BetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessNumberBet, err = convertStringToFloat64(guessNumberBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessOneToEighteenBet, err = convertStringToFloat64(guessOneToEighteenBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessNineteenToThirtySixBet, err = convertStringToFloat64(guessNineteenToThirtySixBetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessFirst2To1Bet, err = convertStringToFloat64(guessFirst2To1BetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessSecond2To1Bet, err = convertStringToFloat64(guessSecond2To1BetStr)
    if err != nil {
        return GameParams{Err: err}
    }
    guessThird2To1Bet, err = convertStringToFloat64(guessThird2To1BetStr)
    if err != nil {
        return GameParams{Err: err}
    }

    // Возвращаем структуру GameParams с заполненными значениями
    return GameParams{
        GuessEvenBet:             guessEvenBet,
        GuessOddBet:              guessOddBet,
        GuessRedBet:              guessRedBet,
        GuessBlackBet:            guessBlackBet,
        GuessSector1st12Bet:      guessSector1st12Bet,
        GuessSector2nd12Bet:      guessSector2nd12Bet,
        GuessSector3rd12Bet:      guessSector3rd12Bet,
        GuessNumberBet:           guessNumberBet,
        GuessOneToEighteenBet:    guessOneToEighteenBet,
        GuessNineteenToThirtySix: guessNineteenToThirtySixBet,
        GuessFirst2To1Bet:        guessFirst2To1Bet,
        GuessSecond2To1Bet:       guessSecond2To1Bet,
        GuessThird2To1Bet:        guessThird2To1Bet,
        Err:                      nil,
    }
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

    // Получаем структуру с игровыми параметрами
    gameParams := GetGameParams(c)

    // Делаем ключи и прокидываем ставку для того чтобы передать в NormalPlay и UnFairPlay
    evenToBets := make(map[string]float64)
    oddToBets := make(map[string]float64)
    redToBets := make(map[string]float64)
    blackToBets := make(map[string]float64)
    sectorsToBets := make(map[string]float64)

    evenToBets["even"] = gameParams.GuessEvenBet
    oddToBets["odd"] = gameParams.GuessOddBet
    redToBets["red"] = gameParams.GuessRedBet
    blackToBets["black"] = gameParams.GuessBlackBet
    
    if gameParams.GuessSector1st12Bet > 0{
        sectorsToBets["1 st 12"] = gameParams.GuessSector1st12Bet
    }
    if gameParams.GuessSector2nd12Bet > 0{
        sectorsToBets["2 nd 12"] = gameParams.GuessSector2nd12Bet
    }
    if gameParams.GuessSector3rd12Bet > 0{
        sectorsToBets["3 rd 12"] = gameParams.GuessSector3rd12Bet
    }

    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game parameters"})
        return
    }

    if fairPlay {
        user_player.NormalPlay(evenToBets, guessNumberInt, betFloat, gameName)
    } else {
        user_player.UnFairPlay(evenToBets, guessNumberInt, betFloat, gameName)
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