package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/services"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
)


type RegisterInput struct{
	Name string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}

type UserInput struct {
	Name string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func RegisterHandler(c *gin.Context){
	var input RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

	hashed_password, err := services.HashPassword(input.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Name = input.Name
	u.Password = hashed_password
	u.Email = input.Email
	u.Balance = input.Balance

	err = userRepository.CreateUser(&u)
    if err != nil {
        // If there's an error creating the user, panic
        panic(err)
    }

}


func LoginHandler(c *gin.Context) {
	var userInput UserInput
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
	if services.CheckPasswordHash(userInput.Password, user.Password){
		tokenString, err := services.CreateToken(userInput.Name)
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