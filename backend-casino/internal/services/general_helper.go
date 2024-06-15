package services

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
    fmt.Println("Dropped 55")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret-key"), nil
    })
    fmt.Println("Dropped 60")
    // Тут дропается код
    if err != nil || !token.Valid {
        return "", errors.New("Invalid token")
    }
    fmt.Println("Dropped 64")
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("Invalid token claims")
    }
    fmt.Println("Dropped 69")
    username, ok := claims["username"].(string)
    fmt.Println(username)
    if !ok {
        return "", errors.New("Invalid user ID in token")
    }
    fmt.Println("Dropped 74")

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
    guessNumberBet, err = convertStringToFloat64(guessNumberBetStr)
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


type GameParamsV2 struct {
    GuessEvenBet             float64
    GuessOddBet              float64
    GuessRedBet              float64
    GuessBlackBet            float64
    GuessSector1st12Bet      float64
    GuessSector2nd12Bet      float64
    GuessSector3rd12Bet      float64
    GuessNumberBet           [37]float64
    GuessNumber              [37]int
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
    NumberToBets           map[[37]int][37]float64
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
        NumberToBets:            make(map[[37]int][37]float64),
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
    // Fill array -1
    for key, _ := range betMaps.NumberToBets{
        for i := 0; i < 37; i++{
            key[i] = -1
        }
    }

    // Count range to copy
    counter := 0
    for i := 0; i < 37 - 1; i++{
        if gameParams.GuessNumber[i] == 0 && gameParams.GuessNumber[i + 1] == 0{
            break
        }
        counter += 1
    }
    // {1, 2, 3, 0, 0, 0}
    // i = 0; counter = 1
    // i = 1; counter = 2
    // i = 2; counter = 3
    
    // Copy normal number to key of betMaps.NumberToBets
    for key, _ := range betMaps.NumberToBets{
        for i := 0; i < counter; i++{
            key[i] = gameParams.GuessNumber[i]
        }
    }

    // Get the key(stupid way)
    var resultKey [37]int
    for key := range betMaps.NumberToBets{
        resultKey = key
        break
    }

    betMaps.NumberToBets[resultKey] = gameParams.GuessNumberBet
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