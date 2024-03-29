package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func SpinRouletteHandler(c *gin.Context) {
    // Получаем JWT токен из заголовка запроса
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
        return
    }

    // Извлекаем токен из заголовка, пропуская префикс "Bearer "
    tokenString := strings.Split(authHeader, " ")[1]

    // Проверяем и парсим токен
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Здесь нужно вернуть ключ подписи токена, используемый при подписи JWT
        return []byte("your_secret_key"), nil
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    // Извлекаем идентификатор пользователя из токена
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        return
    }
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
        return
    }
    userID := uint(userIDFloat)

    // Продолжаем обработку запроса, используя полученного пользователя
    db := database.InitDB()
    user_repository := repositories.UserRepository{Db: db}

    user, err := user_repository.GetUserById(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    user_player := services.UserPlayer{}
    user_player.Balance = user.Balance

    guess_number := c.PostForm("guess_number")
    guessNumberToInt, err := strconv.Atoi(guess_number)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guess number"})
        return
    }

    bet := c.PostForm("bet")
    betToInt, err := strconv.Atoi(bet)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet"})
        return
    }

    gameName := c.PostForm("gameName")
    user_player.Play(guessNumberToInt, betToInt, gameName)
}




func CreateGameHandler(c *gin.Context) {
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


func GetGameByIdHandler(c *gin.Context) {
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


func UpdateGameHandler(c *gin.Context) {
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


func DeleteGameHandler(c *gin.Context) {
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
