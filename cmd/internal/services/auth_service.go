package services

import (
	"crypto/sha256"
	"fmt"
	"time"
    "encoding/hex"

	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/golang-jwt/jwt"

)


var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}


func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return secretKey, nil
	})
   
	if err != nil {
	   return err
	}
   
	if !token.Valid {
	   return fmt.Errorf("invalid token")
	}
   
	return nil
 }


 func HashPassword(password string) string {
    hasher := sha256.New()
    hasher.Write([]byte(password))
    hash := hasher.Sum(nil)
    return hex.EncodeToString(hash)
}

func CheckPasswordHash(password, hash string) bool {
    inputHash := HashPassword(password)
    return inputHash == hash
}
type RegisterInput struct{
	Name string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}

type UserInput struct {
	Name string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func CreateUser(input RegisterInput) error {
    db := database.InitDB()


    userRepository := repositories.UserRepository{Db: db}

    hashedPassword := HashPassword(input.Password)

    u := models.User{
        Name:     input.Name,
        Password: hashedPassword,
        Email:    input.Email,
        Balance:  input.Balance,
    }

    if err := userRepository.CreateUser(&u); err != nil {
        return err
    }

    return nil
}