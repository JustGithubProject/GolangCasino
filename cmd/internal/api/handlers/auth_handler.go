package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/services"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"net/http"
	"github.com/gin-gonic/gin"
)


type RegisterInput struct{
	Name string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
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

	u := models.User{}
	u.Name = input.Name
	u.Password = input.Password
	u.Email = input.Email
	u.Balance = input.Balance

	err = userRepository.CreateUser(&u)
    if err != nil {
        // If there's an error creating the user, panic
        panic(err)
    }

}


func LoginHandler(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if u.Name == "Chek" && u.Password == "123456" {
		tokenString, err := services.CreateToken(u.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}