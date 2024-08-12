package services

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
    "log"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/database"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func ShuffleWeights(arr []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}


func InitWeights(length int) []int{
    weights_arr := make([]int, 37)
	counter_weight := 10

	for i := 0; i < length; i++{
		weights_arr[i] = counter_weight
		counter_weight += 100
	}
    return weights_arr
}


func InitNumbersArray() []int{
    arr := make([]int, 37)
    for i := 0; i < 37; i++ {
        arr[i] = i
    }
    return arr
}



func ValidateToken(c *gin.Context) (string, error) {
    authHeader := c.GetHeader("Authorization")

    if authHeader == "" {
        return "", errors.New("Authorization header is missing")
    }
    
    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) < 2 {
        return "", errors.New("Invalid Authorization header format")
    }
    
    tokenString := tokenParts[1]
    log.Println("Auth token: ", tokenString)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret-key"), nil
    })
    // Тут дропается код
    if err != nil || !token.Valid {
        return "", errors.New("Invalid token")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("Invalid token claims")
    }

    username, ok := claims["username"].(string)
    if !ok {
        return "", errors.New("Invalid user ID in token")
    }

    return username, nil
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
    guessEvenBetStr := c.Query("even")
    guessOddBetStr := c.Query("odd")
    guessRedBetStr := c.Query("red")
    guessBlackBetStr := c.Query("black")
    guessSector1st12BetStr := c.Query("1st12")
    guessSector2nd12BetStr := c.Query("2nd12")
    guessSector3rd12BetStr := c.Query("3rd12")
    guessNumberBetStr := c.Query("number")
    guessNumberStr := c.Query("num")
    guessOneToEighteenBetStr := c.Query("1To18")
    guessNineteenToThirtySixBetStr := c.Query("19To36")
    guessFirst2To1BetStr := c.Query("First2To1")
    guessSecond2To1BetStr := c.Query("Second2To1")
    guessThird2To1BetStr := c.Query("Third2To1")


    // Функция для преобразования строки в float64
    convertStringToFloat64 := func(str string) (float64, error) {
        if str == "" {
            return 0, nil // Возвращаем 0, если строка пустая
        }
        return strconv.ParseFloat(str, 64)
    }

    // Преобразование и обработка значений

    guessNumberInt, _ := strconv.Atoi(guessNumberStr)

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


// Функция для получения параметров игры из запроса
func GetGameParamsV2(c *gin.Context) GameParamsV2 {
    var (
        guessEvenBet             float64
        guessOddBet              float64
        guessRedBet              float64
        guessBlackBet            float64
        guessSector1st12Bet      float64
        guessSector2nd12Bet      float64
        guessSector3rd12Bet      float64
        guessOneToEighteenBet    float64
        guessNineteenToThirtySixBet float64
        guessFirst2To1Bet        float64
        guessSecond2To1Bet       float64
        guessThird2To1Bet        float64
        err                      error
    )

    guessEvenBetStr := c.Query("even")
    guessOddBetStr := c.Query("odd")
    guessRedBetStr := c.Query("red")
    guessBlackBetStr := c.Query("black")
    guessSector1st12BetStr := c.Query("1st12")
    guessSector2nd12BetStr := c.Query("2nd12")
    guessSector3rd12BetStr := c.Query("3rd12")
    // guessNumberBetStr := c.Query("number")
    guessNumberCell_0 := c.Query("num_0") // ?num_0=82.122
    guessNumberCell_1 := c.Query("num_1") 
    guessNumberCell_2 := c.Query("num_2")
    guessNumberCell_3 := c.Query("num_3")
    guessNumberCell_4 := c.Query("num_4")
    guessNumberCell_5 := c.Query("num_5")
    guessNumberCell_6 := c.Query("num_6")
    guessNumberCell_7 := c.Query("num_7")
    guessNumberCell_8 := c.Query("num_8")
    guessNumberCell_9 := c.Query("num_9")
    guessNumberCell_10 := c.Query("num_10")
    guessNumberCell_11 := c.Query("num_11")
    guessNumberCell_12 := c.Query("num_12")
    guessNumberCell_13 := c.Query("num_13")
    guessNumberCell_14 := c.Query("num_14")
    guessNumberCell_15 := c.Query("num_15")
    guessNumberCell_16 := c.Query("num_16")
    guessNumberCell_17 := c.Query("num_17")
    guessNumberCell_18 := c.Query("num_18")
    guessNumberCell_19 := c.Query("num_19")
    guessNumberCell_20 := c.Query("num_20")
    guessNumberCell_21 := c.Query("num_21")
    guessNumberCell_22 := c.Query("num_22")
    guessNumberCell_23 := c.Query("num_23")
    guessNumberCell_24 := c.Query("num_24")
    guessNumberCell_25 := c.Query("num_25")
    guessNumberCell_26 := c.Query("num_26")
    guessNumberCell_27 := c.Query("num_27")
    guessNumberCell_28 := c.Query("num_28")
    guessNumberCell_29 := c.Query("num_29")
    guessNumberCell_30 := c.Query("num_30")
    guessNumberCell_31 := c.Query("num_31")
    guessNumberCell_32 := c.Query("num_32")
    guessNumberCell_33 := c.Query("num_33")
    guessNumberCell_34 := c.Query("num_34")
    guessNumberCell_35 := c.Query("num_35")
    guessNumberCell_36 := c.Query("num_36")
    guessOneToEighteenBetStr := c.Query("1To18")
    guessNineteenToThirtySixBetStr := c.Query("19To36")
    guessFirst2To1BetStr := c.Query("First2To1")
    guessSecond2To1BetStr := c.Query("Second2To1")
    guessThird2To1BetStr := c.Query("Third2To1")


    // Функция для преобразования строки в float64
    convertStringToFloat64 := func(str string) (float64, error) {
        if str == "" {
            return 0, nil // Возвращаем 0, если строка пустая
        }
        return strconv.ParseFloat(str, 64)
    }

    // Преобразование и обработка значений
    guessNumberCell_0BET, err := convertStringToFloat64(guessNumberCell_0)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_1BET, err := convertStringToFloat64(guessNumberCell_1)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_2BET, err := convertStringToFloat64(guessNumberCell_2)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_3BET, err := convertStringToFloat64(guessNumberCell_3)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_4BET, err := convertStringToFloat64(guessNumberCell_4)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_5BET, err := convertStringToFloat64(guessNumberCell_5)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_6BET, err := convertStringToFloat64(guessNumberCell_6)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_7BET, err := convertStringToFloat64(guessNumberCell_7)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_8BET, err := convertStringToFloat64(guessNumberCell_8)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_9BET, err := convertStringToFloat64(guessNumberCell_9)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_10BET, err := convertStringToFloat64(guessNumberCell_10)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_11BET, err := convertStringToFloat64(guessNumberCell_11)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_12BET, err := convertStringToFloat64(guessNumberCell_12)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_13BET, err := convertStringToFloat64(guessNumberCell_13)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_14BET, err := convertStringToFloat64(guessNumberCell_14)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_15BET, err := convertStringToFloat64(guessNumberCell_15)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_16BET, err := convertStringToFloat64(guessNumberCell_16)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_17BET, err := convertStringToFloat64(guessNumberCell_17)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_18BET, err := convertStringToFloat64(guessNumberCell_18)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_19BET, err := convertStringToFloat64(guessNumberCell_19)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_20BET, err := convertStringToFloat64(guessNumberCell_20)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_21BET, err := convertStringToFloat64(guessNumberCell_21)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_22BET, err := convertStringToFloat64(guessNumberCell_22)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_23BET, err := convertStringToFloat64(guessNumberCell_23)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_24BET, err := convertStringToFloat64(guessNumberCell_24)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_25BET, err := convertStringToFloat64(guessNumberCell_25)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_26BET, err := convertStringToFloat64(guessNumberCell_26)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_27BET, err := convertStringToFloat64(guessNumberCell_27)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_28BET, err := convertStringToFloat64(guessNumberCell_28)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_29BET, err := convertStringToFloat64(guessNumberCell_29)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_30BET, err := convertStringToFloat64(guessNumberCell_30)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_31BET, err := convertStringToFloat64(guessNumberCell_31)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_32BET, err := convertStringToFloat64(guessNumberCell_32)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_33BET, err := convertStringToFloat64(guessNumberCell_33)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_34BET, err := convertStringToFloat64(guessNumberCell_34)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_35BET, err := convertStringToFloat64(guessNumberCell_35)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNumberCell_36BET, err := convertStringToFloat64(guessNumberCell_36)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    

    guessEvenBet, err = convertStringToFloat64(guessEvenBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessOddBet, err = convertStringToFloat64(guessOddBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessRedBet, err = convertStringToFloat64(guessRedBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessBlackBet, err = convertStringToFloat64(guessBlackBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessSector1st12Bet, err = convertStringToFloat64(guessSector1st12BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessSector2nd12Bet, err = convertStringToFloat64(guessSector2nd12BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessSector3rd12Bet, err = convertStringToFloat64(guessSector3rd12BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessOneToEighteenBet, err = convertStringToFloat64(guessOneToEighteenBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessNineteenToThirtySixBet, err = convertStringToFloat64(guessNineteenToThirtySixBetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessFirst2To1Bet, err = convertStringToFloat64(guessFirst2To1BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessSecond2To1Bet, err = convertStringToFloat64(guessSecond2To1BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }
    guessThird2To1Bet, err = convertStringToFloat64(guessThird2To1BetStr)
    if err != nil {
        return GameParamsV2{Err: err}
    }

    // Возвращаем структуру GameParams с заполненными значениями
    return GameParamsV2{
        GuessEvenBet:             guessEvenBet,
        GuessOddBet:              guessOddBet,
        GuessRedBet:              guessRedBet,
        GuessBlackBet:            guessBlackBet,
        GuessSector1st12Bet:      guessSector1st12Bet,
        GuessSector2nd12Bet:      guessSector2nd12Bet,
        GuessSector3rd12Bet:      guessSector3rd12Bet,

        GuessNumberCell_0Bet:     guessNumberCell_0BET,
        GuessNumberCell_1Bet:     guessNumberCell_1BET,
        GuessNumberCell_2Bet:     guessNumberCell_2BET,
        GuessNumberCell_3Bet:     guessNumberCell_3BET,
        GuessNumberCell_4Bet:     guessNumberCell_4BET,
        GuessNumberCell_5Bet:     guessNumberCell_5BET,
        GuessNumberCell_6Bet:     guessNumberCell_6BET,
        GuessNumberCell_7Bet:     guessNumberCell_7BET,
        GuessNumberCell_8Bet:     guessNumberCell_8BET,
        GuessNumberCell_9Bet:     guessNumberCell_9BET,
        GuessNumberCell_10Bet:    guessNumberCell_10BET,
        GuessNumberCell_11Bet:    guessNumberCell_11BET,
        GuessNumberCell_12Bet:    guessNumberCell_12BET,
        GuessNumberCell_13Bet:    guessNumberCell_13BET,
        GuessNumberCell_14Bet:    guessNumberCell_14BET,
        GuessNumberCell_15Bet:    guessNumberCell_15BET,
        GuessNumberCell_16Bet:    guessNumberCell_16BET,
        GuessNumberCell_17Bet:    guessNumberCell_17BET,
        GuessNumberCell_18Bet:    guessNumberCell_18BET,
        GuessNumberCell_19Bet:    guessNumberCell_19BET,
        GuessNumberCell_20Bet:    guessNumberCell_20BET,
        GuessNumberCell_21Bet:    guessNumberCell_21BET,
        GuessNumberCell_22Bet:    guessNumberCell_22BET,
        GuessNumberCell_23Bet:    guessNumberCell_23BET,
        GuessNumberCell_24Bet:    guessNumberCell_24BET,
        GuessNumberCell_25Bet:    guessNumberCell_25BET,
        GuessNumberCell_26Bet:    guessNumberCell_26BET,
        GuessNumberCell_27Bet:    guessNumberCell_27BET,
        GuessNumberCell_28Bet:    guessNumberCell_28BET,
        GuessNumberCell_29Bet:    guessNumberCell_29BET,
        GuessNumberCell_30Bet:    guessNumberCell_30BET,
        GuessNumberCell_31Bet:    guessNumberCell_31BET,
        GuessNumberCell_32Bet:    guessNumberCell_32BET,
        GuessNumberCell_33Bet:    guessNumberCell_33BET,
        GuessNumberCell_34Bet:    guessNumberCell_34BET,
        GuessNumberCell_35Bet:    guessNumberCell_35BET,
        GuessNumberCell_36Bet:    guessNumberCell_36BET,
         
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


type GameParamsV2 struct {
    GuessEvenBet             float64
    GuessOddBet              float64
    GuessRedBet              float64
    GuessBlackBet            float64
    GuessSector1st12Bet      float64
    GuessSector2nd12Bet      float64
    GuessSector3rd12Bet      float64
    GuessNumberCell_0Bet     float64
    GuessNumberCell_1Bet     float64
    GuessNumberCell_2Bet     float64
    GuessNumberCell_3Bet     float64
    GuessNumberCell_4Bet     float64
    GuessNumberCell_5Bet     float64
    GuessNumberCell_6Bet     float64
    GuessNumberCell_7Bet     float64
    GuessNumberCell_8Bet     float64
    GuessNumberCell_9Bet     float64
    GuessNumberCell_10Bet     float64
    GuessNumberCell_11Bet     float64
    GuessNumberCell_12Bet     float64
    GuessNumberCell_13Bet     float64
    GuessNumberCell_14Bet     float64
    GuessNumberCell_15Bet     float64
    GuessNumberCell_16Bet     float64
    GuessNumberCell_17Bet     float64
    GuessNumberCell_18Bet     float64
    GuessNumberCell_19Bet     float64
    GuessNumberCell_20Bet     float64
    GuessNumberCell_21Bet     float64
    GuessNumberCell_22Bet     float64
    GuessNumberCell_23Bet     float64
    GuessNumberCell_24Bet     float64
    GuessNumberCell_25Bet     float64
    GuessNumberCell_26Bet     float64
    GuessNumberCell_27Bet     float64
    GuessNumberCell_28Bet     float64
    GuessNumberCell_29Bet     float64
    GuessNumberCell_30Bet     float64
    GuessNumberCell_31Bet     float64
    GuessNumberCell_32Bet     float64
    GuessNumberCell_33Bet     float64
    GuessNumberCell_34Bet     float64
    GuessNumberCell_35Bet     float64
    GuessNumberCell_36Bet     float64
    GuessOneToEighteenBet    float64
    GuessNineteenToThirtySix float64
    GuessFirst2To1Bet        float64
    GuessSecond2To1Bet       float64
    GuessThird2To1Bet        float64
    Err                      error
}



type BetMaps struct {
    EvenToBets             map[string]float64
    OddToBets              map[string]float64
    RedToBets              map[string]float64
    BlackToBets            map[string]float64
    SectorsToBets          map[string]float64
    NumberToBets           map[int]float64
    OneToEighteenBets      map[string]float64
    NineteenToThirtySixBets map[string]float64
    First2To1Bets          map[string]float64
    Second2To1Bets         map[string]float64
    Third2To1Bets          map[string]float64
}


type BetMapsV2 struct {
    EvenToBets             map[string]float64
    OddToBets              map[string]float64
    RedToBets              map[string]float64
    BlackToBets            map[string]float64
    SectorsToBets          map[string]float64
    NumberCell_0           map[int]float64
    NumberCell_1           map[int]float64
    NumberCell_2           map[int]float64
    NumberCell_3           map[int]float64
    NumberCell_4           map[int]float64
    NumberCell_5           map[int]float64
    NumberCell_6           map[int]float64
    NumberCell_7           map[int]float64
    NumberCell_8           map[int]float64
    NumberCell_9           map[int]float64
    NumberCell_10          map[int]float64
    NumberCell_11          map[int]float64
    NumberCell_12          map[int]float64
    NumberCell_13          map[int]float64
    NumberCell_14          map[int]float64
    NumberCell_15          map[int]float64
    NumberCell_16          map[int]float64
    NumberCell_17          map[int]float64
    NumberCell_18          map[int]float64
    NumberCell_19          map[int]float64
    NumberCell_20          map[int]float64
    NumberCell_21          map[int]float64
    NumberCell_22          map[int]float64
    NumberCell_23          map[int]float64
    NumberCell_24          map[int]float64
    NumberCell_25          map[int]float64
    NumberCell_26          map[int]float64
    NumberCell_27          map[int]float64
    NumberCell_28          map[int]float64
    NumberCell_29          map[int]float64
    NumberCell_30          map[int]float64
    NumberCell_31          map[int]float64
    NumberCell_32          map[int]float64
    NumberCell_33          map[int]float64
    NumberCell_34          map[int]float64
    NumberCell_35          map[int]float64
    NumberCell_36          map[int]float64
    OneToEighteenBets      map[string]float64
    NineteenToThirtySixBets map[string]float64
    First2To1Bets          map[string]float64
    Second2To1Bets         map[string]float64
    Third2To1Bets          map[string]float64
}


func InitBetsMapV2(gameParams GameParamsV2) BetMapsV2 {
    betMaps := BetMapsV2{
        EvenToBets:              make(map[string]float64),
        OddToBets:               make(map[string]float64),
        RedToBets:               make(map[string]float64),
        BlackToBets:             make(map[string]float64),
        SectorsToBets:           make(map[string]float64),
        NumberCell_0:            make(map[int]float64),
        NumberCell_1:            make(map[int]float64),
        NumberCell_2:            make(map[int]float64),
        NumberCell_3:            make(map[int]float64),
        NumberCell_4:            make(map[int]float64),
        NumberCell_5:            make(map[int]float64),
        NumberCell_6:            make(map[int]float64),
        NumberCell_7:            make(map[int]float64),
        NumberCell_8:            make(map[int]float64),
        NumberCell_9:            make(map[int]float64),
        NumberCell_10:           make(map[int]float64),
        NumberCell_11:           make(map[int]float64),
        NumberCell_12:           make(map[int]float64),
        NumberCell_13:           make(map[int]float64),
        NumberCell_14:           make(map[int]float64),
        NumberCell_15:           make(map[int]float64),
        NumberCell_16:           make(map[int]float64),
        NumberCell_17:           make(map[int]float64),
        NumberCell_18:           make(map[int]float64),
        NumberCell_19:           make(map[int]float64),
        NumberCell_20:           make(map[int]float64),
        NumberCell_21:           make(map[int]float64),
        NumberCell_22:           make(map[int]float64),
        NumberCell_23:           make(map[int]float64),
        NumberCell_24:           make(map[int]float64),
        NumberCell_25:           make(map[int]float64),
        NumberCell_26:           make(map[int]float64),
        NumberCell_27:           make(map[int]float64),
        NumberCell_28:           make(map[int]float64),
        NumberCell_29:           make(map[int]float64),
        NumberCell_30:           make(map[int]float64),
        NumberCell_31:           make(map[int]float64),
        NumberCell_32:           make(map[int]float64),
        NumberCell_33:           make(map[int]float64),
        NumberCell_34:           make(map[int]float64),
        NumberCell_35:           make(map[int]float64),
        NumberCell_36:           make(map[int]float64),
        OneToEighteenBets:       make(map[string]float64),
        NineteenToThirtySixBets: make(map[string]float64),
        First2To1Bets:           make(map[string]float64),
        Second2To1Bets:          make(map[string]float64),
        Third2To1Bets:           make(map[string]float64),
    }

    betMaps.EvenToBets["even"] = gameParams.GuessEvenBet
    betMaps.OddToBets["odd"] = gameParams.GuessOddBet
    betMaps.RedToBets["red"] = gameParams.GuessRedBet
    betMaps.BlackToBets["black"] = gameParams.GuessBlackBet

    // Sectors (1 st 12, 2 nd 12, 3 rd 12. КОСТЫЛЬ)
    if gameParams.GuessSector1st12Bet > 0 {
        betMaps.SectorsToBets["1 st 12"] = gameParams.GuessSector1st12Bet
    }
    if gameParams.GuessSector2nd12Bet > 0 {
        betMaps.SectorsToBets["2 nd 12"] = gameParams.GuessSector2nd12Bet
    }
    if gameParams.GuessSector3rd12Bet > 0 {
        betMaps.SectorsToBets["3 rd 12"] = gameParams.GuessSector3rd12Bet
    }

    betMaps.NumberCell_0[0] = gameParams.GuessNumberCell_0Bet
    betMaps.NumberCell_1[1] = gameParams.GuessNumberCell_1Bet
    betMaps.NumberCell_2[2] = gameParams.GuessNumberCell_2Bet
    betMaps.NumberCell_3[3] = gameParams.GuessNumberCell_3Bet
    betMaps.NumberCell_4[4] = gameParams.GuessNumberCell_4Bet
    betMaps.NumberCell_5[5] = gameParams.GuessNumberCell_5Bet
    betMaps.NumberCell_6[6] = gameParams.GuessNumberCell_6Bet
    betMaps.NumberCell_7[7] = gameParams.GuessNumberCell_7Bet
    betMaps.NumberCell_8[8] = gameParams.GuessNumberCell_8Bet
    betMaps.NumberCell_9[9] = gameParams.GuessNumberCell_9Bet
    betMaps.NumberCell_10[10] = gameParams.GuessNumberCell_10Bet
    betMaps.NumberCell_11[11] = gameParams.GuessNumberCell_11Bet
    betMaps.NumberCell_12[12] = gameParams.GuessNumberCell_12Bet
    betMaps.NumberCell_13[13] = gameParams.GuessNumberCell_13Bet
    betMaps.NumberCell_14[14] = gameParams.GuessNumberCell_14Bet
    betMaps.NumberCell_15[15] = gameParams.GuessNumberCell_15Bet
    betMaps.NumberCell_16[16] = gameParams.GuessNumberCell_16Bet
    betMaps.NumberCell_17[17] = gameParams.GuessNumberCell_17Bet
    betMaps.NumberCell_18[18] = gameParams.GuessNumberCell_18Bet
    betMaps.NumberCell_19[19] = gameParams.GuessNumberCell_19Bet
    betMaps.NumberCell_20[20] = gameParams.GuessNumberCell_20Bet
    betMaps.NumberCell_21[21] = gameParams.GuessNumberCell_21Bet
    betMaps.NumberCell_22[22] = gameParams.GuessNumberCell_22Bet
    betMaps.NumberCell_23[23] = gameParams.GuessNumberCell_23Bet
    betMaps.NumberCell_24[24] = gameParams.GuessNumberCell_24Bet
    betMaps.NumberCell_25[25] = gameParams.GuessNumberCell_25Bet
    betMaps.NumberCell_26[26] = gameParams.GuessNumberCell_26Bet
    betMaps.NumberCell_27[27] = gameParams.GuessNumberCell_27Bet
    betMaps.NumberCell_28[28] = gameParams.GuessNumberCell_28Bet
    betMaps.NumberCell_29[29] = gameParams.GuessNumberCell_29Bet
    betMaps.NumberCell_30[30] = gameParams.GuessNumberCell_30Bet
    betMaps.NumberCell_31[31] = gameParams.GuessNumberCell_31Bet
    betMaps.NumberCell_32[32] = gameParams.GuessNumberCell_32Bet
    betMaps.NumberCell_33[33] = gameParams.GuessNumberCell_33Bet
    betMaps.NumberCell_34[34] = gameParams.GuessNumberCell_34Bet
    betMaps.NumberCell_35[35] = gameParams.GuessNumberCell_35Bet
    betMaps.NumberCell_36[36] = gameParams.GuessNumberCell_36Bet

    betMaps.OneToEighteenBets["1to18"] = gameParams.GuessOneToEighteenBet
    betMaps.NineteenToThirtySixBets["19to36"] = gameParams.GuessNineteenToThirtySix
    betMaps.First2To1Bets["2to1"] = gameParams.GuessFirst2To1Bet
    betMaps.Second2To1Bets["2to1"] = gameParams.GuessSecond2To1Bet
    betMaps.Third2To1Bets["2to1"] = gameParams.GuessThird2To1Bet

    return betMaps
}



func InitBetsMap(gameParams GameParams) BetMaps {
    betMaps := BetMaps{
        EvenToBets:              make(map[string]float64),
        OddToBets:               make(map[string]float64),
        RedToBets:               make(map[string]float64),
        BlackToBets:             make(map[string]float64),
        SectorsToBets:           make(map[string]float64),
        NumberToBets:            make(map[int]float64),
        OneToEighteenBets:       make(map[string]float64),
        NineteenToThirtySixBets: make(map[string]float64),
        First2To1Bets:           make(map[string]float64),
        Second2To1Bets:          make(map[string]float64),
        Third2To1Bets:           make(map[string]float64),
    }

    betMaps.EvenToBets["even"] = gameParams.GuessEvenBet
    betMaps.OddToBets["odd"] = gameParams.GuessOddBet
    betMaps.RedToBets["red"] = gameParams.GuessRedBet
    betMaps.BlackToBets["black"] = gameParams.GuessBlackBet

    // Sectors (1 st 12, 2 nd 12, 3 rd 12. КОСТЫЛЬ)
    if gameParams.GuessSector1st12Bet > 0 {
        betMaps.SectorsToBets["1 st 12"] = gameParams.GuessSector1st12Bet
    }
    if gameParams.GuessSector2nd12Bet > 0 {
        betMaps.SectorsToBets["2 nd 12"] = gameParams.GuessSector2nd12Bet
    }
    if gameParams.GuessSector3rd12Bet > 0 {
        betMaps.SectorsToBets["3 rd 12"] = gameParams.GuessSector3rd12Bet
    }

    betMaps.NumberToBets[gameParams.GuessNumber] = gameParams.GuessNumberBet
    betMaps.OneToEighteenBets["1to18"] = gameParams.GuessOneToEighteenBet
    betMaps.NineteenToThirtySixBets["19to36"] = gameParams.GuessNineteenToThirtySix
    betMaps.First2To1Bets["2to1"] = gameParams.GuessFirst2To1Bet
    betMaps.Second2To1Bets["2to1"] = gameParams.GuessSecond2To1Bet
    betMaps.Third2To1Bets["2to1"] = gameParams.GuessThird2To1Bet

    return betMaps
}