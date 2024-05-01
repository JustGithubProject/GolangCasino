package services

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
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
    guessNumberStr := c.PostForm("num")
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

    guessNumberInt, err := strconv.Atoi(guessNumberStr)

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
        GuessNumber:              guessNumberInt,  
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

type GameParams struct {
    GuessEvenBet             float64
    GuessOddBet              float64
    GuessRedBet              float64
    GuessBlackBet            float64
    GuessSector1st12Bet      float64
    GuessSector2nd12Bet      float64
    GuessSector3rd12Bet      float64
    GuessNumberBet           float64
    GuessNumber              int
    GuessOneToEighteenBet    float64
    GuessNineteenToThirtySix float64
    GuessFirst2To1Bet        float64
    GuessSecond2To1Bet       float64
    GuessThird2To1Bet        float64
    Err                      error
}