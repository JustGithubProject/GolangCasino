package services

import (
	"database/sql"
	"fmt"
	"strings"

	"net/http"
	"strconv"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
	"github.com/gin-gonic/gin"
)




func HandleGameRequest(c *gin.Context, fairPlay bool) {
    // Получаем JWT токен из заголовка запроса
    username, err := ValidateToken(c)
    if err != nil {
        fmt.Println("С токеном проблемы?")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    user_repository, err := InitializeUserRepository()
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

    user_player := UserPlayer{}
    user_player.Balance = user.Balance

    // Получаем структуру с игровыми параметрами
    gameParams := GetGameParams(c)
    fmt.Printf("GameParams: %v+\n", gameParams)
    // Делаем ключи и прокидываем ставку для того чтобы передать в NormalPlay и UnFairPlay
    betMaps := InitBetsMap(gameParams)
    fmt.Printf("BetMaps: %v+\n", betMaps)
    if err != nil {
        fmt.Println("Failure to get game parameters")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game parameters"})
        return
    }
    fmt.Printf("fairPlay: %t", fairPlay)
    currentBalance := user.Balance
    totalBet := user_player.getTotalBet(
        betMaps.EvenToBets,
        betMaps.OddToBets,
        betMaps.RedToBets,
        betMaps.BlackToBets,
        betMaps.SectorsToBets,
        betMaps.NumberToBets,
        betMaps.OneToEighteenBets,
        betMaps.NineteenToThirtySixBets,
        betMaps.First2To1Bets,
        betMaps.Second2To1Bets,
        betMaps.Third2To1Bets,
	)
    if totalBet > currentBalance{
        fmt.Println("You don't have enough funds")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place a bet with insufficient funds"})
        return
    }
    var dropped_number int
    if fairPlay {
        currentBalance, dropped_number, err = user_player.NormalPlay(
            betMaps.EvenToBets,
            betMaps.OddToBets,
            betMaps.RedToBets,
            betMaps.BlackToBets,
            betMaps.SectorsToBets,
            betMaps.NumberToBets,
            betMaps.OneToEighteenBets,
            betMaps.NineteenToThirtySixBets,
            betMaps.First2To1Bets,
            betMaps.Second2To1Bets,
            betMaps.Third2To1Bets,
        )
    } else {
        currentBalance, dropped_number, err = user_player.UnFairPlay(
            betMaps.EvenToBets,
            betMaps.OddToBets,
            betMaps.RedToBets,
            betMaps.BlackToBets,
            betMaps.SectorsToBets,
            betMaps.NumberToBets,
            betMaps.OneToEighteenBets,
            betMaps.NineteenToThirtySixBets,
            betMaps.First2To1Bets,
            betMaps.Second2To1Bets,
            betMaps.Third2To1Bets,
        )
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user.Balance = currentBalance
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Game request handled successfully", "user": user, "dropped_number": dropped_number})
}

func HandleVeryBadGameRequest(c *gin.Context){
    username, err := ValidateToken(c)
    if err != nil {
        fmt.Println("Проблема с токеном?")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
        return
    }
    
    user_repository, err := InitializeUserRepository()
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

    user_player := UserPlayer{}
    user_player.Balance = user.Balance

    // Получаем структуру с игровыми параметрами
    gameParams := GetGameParamsV2(c)
    fmt.Printf("GameParams: %v+\n", gameParams)
    
    // Делаем ключи и прокидываем ставку для того чтобы передать в VeryBadPlay
    betMaps := InitBetsMapV2(gameParams)
    fmt.Printf("BetMaps: %v+\n", betMaps)
    if err != nil {
        fmt.Println("Failure to get game parameters")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game parameters"})
        return
    }
    currentBalance := user.Balance
    totalBet := user_player.getTotalBetV2(
        betMaps.EvenToBets,
        betMaps.OddToBets,
        betMaps.RedToBets,
        betMaps.BlackToBets,
        betMaps.SectorsToBets,
        betMaps.NumberCell_0,
        betMaps.NumberCell_1,
        betMaps.NumberCell_2,
        betMaps.NumberCell_3,
        betMaps.NumberCell_4,
        betMaps.NumberCell_5,
        betMaps.NumberCell_6,
        betMaps.NumberCell_7,
        betMaps.NumberCell_8,
        betMaps.NumberCell_9,
        betMaps.NumberCell_10,
        betMaps.NumberCell_11,
        betMaps.NumberCell_12,
        betMaps.NumberCell_13,
        betMaps.NumberCell_14,
        betMaps.NumberCell_15,
        betMaps.NumberCell_16,
        betMaps.NumberCell_17,
        betMaps.NumberCell_18,
        betMaps.NumberCell_19,
        betMaps.NumberCell_20,
        betMaps.NumberCell_21,
        betMaps.NumberCell_22,
        betMaps.NumberCell_23,
        betMaps.NumberCell_24,
        betMaps.NumberCell_25,
        betMaps.NumberCell_26,
        betMaps.NumberCell_27,
        betMaps.NumberCell_28,
        betMaps.NumberCell_29,
        betMaps.NumberCell_30,
        betMaps.NumberCell_31,
        betMaps.NumberCell_32,
        betMaps.NumberCell_33,
        betMaps.NumberCell_34,
        betMaps.NumberCell_35,
        betMaps.NumberCell_36,
        betMaps.OneToEighteenBets,
        betMaps.NineteenToThirtySixBets,
        betMaps.First2To1Bets,
        betMaps.Second2To1Bets,
        betMaps.Third2To1Bets,
	)
    if totalBet > currentBalance{
        fmt.Println("You don't have enough funds")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place a bet with insufficient funds"})
        return
    }
    var dropped_number int
    currentBalance, dropped_number, err = user_player.VeryBadPlay(
        betMaps.EvenToBets,
        betMaps.OddToBets,
        betMaps.RedToBets,
        betMaps.BlackToBets,
        betMaps.SectorsToBets,
        betMaps.NumberCell_0,
        betMaps.NumberCell_1,
        betMaps.NumberCell_2,
        betMaps.NumberCell_3,
        betMaps.NumberCell_4,
        betMaps.NumberCell_5,
        betMaps.NumberCell_6,
        betMaps.NumberCell_7,
        betMaps.NumberCell_8,
        betMaps.NumberCell_9,
        betMaps.NumberCell_10,
        betMaps.NumberCell_11,
        betMaps.NumberCell_12,
        betMaps.NumberCell_13,
        betMaps.NumberCell_14,
        betMaps.NumberCell_15,
        betMaps.NumberCell_16,
        betMaps.NumberCell_17,
        betMaps.NumberCell_18,
        betMaps.NumberCell_19,
        betMaps.NumberCell_20,
        betMaps.NumberCell_21,
        betMaps.NumberCell_22,
        betMaps.NumberCell_23,
        betMaps.NumberCell_24,
        betMaps.NumberCell_25,
        betMaps.NumberCell_26,
        betMaps.NumberCell_27,
        betMaps.NumberCell_28,
        betMaps.NumberCell_29,
        betMaps.NumberCell_30,
        betMaps.NumberCell_31,
        betMaps.NumberCell_32,
        betMaps.NumberCell_33,
        betMaps.NumberCell_34,
        betMaps.NumberCell_35,
        betMaps.NumberCell_36,
        betMaps.OneToEighteenBets,
        betMaps.NineteenToThirtySixBets,
        betMaps.First2To1Bets,
        betMaps.Second2To1Bets,
        betMaps.Third2To1Bets,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user.Balance = currentBalance
    err = user_repository.UpdateBalanceUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Game request handled successfully", "user": user, "dropped_number": dropped_number})
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

    // Remove leading and trailing spaces from the password
    input.Password = strings.TrimSpace(input.Password)
    hashedPassword := HashPassword(input.Password)
    input.Password = hashedPassword

    // Creating user
    if err := CreateUser(input); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func HandleUserLogin(c *gin.Context) {
    fmt.Println("Inside of HandleUserLogin")
    var userInput UserInput
    if err := c.BindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "(Bind)Invalid JSON"})
        return
    }

    userRepository, err := InitializeUserRepository()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to init db and repository"})
        return
    }

    user, err := userRepository.GetUserByUsername(userInput.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if CheckPasswordHash(userInput.Password, user.Password) {
        // Check if user has a valid token
        existingToken := c.Request.Header.Get("Authorization")
        fmt.Println("existingToken: " + existingToken)
        if existingToken != "" {
            token, err:= ParseToken(existingToken[7:])
            if err == nil && token.Valid {
                fmt.Println("Валидный?")
                c.JSON(http.StatusOK, gin.H{"token": existingToken})
                return
            }
        }

        // Generate a new token
        fmt.Println("Создаем токен ?")
        tokenString, err := CreateToken(user.Name)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }
        
        // Set the Authorization header with the new token
        fmt.Println("tokenString", tokenString)
        c.Header("Authorization", "Bearer "+tokenString)
        c.JSON(http.StatusOK, gin.H{"token": tokenString})
        return
    }

    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

