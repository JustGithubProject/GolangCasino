package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/gin-gonic/gin"
)



func CreateUserHandler(
	c *gin.Context,
){
	var user models.User
	err_1 := c.BindJSON(&user)
	if err_1 != nil{
		c.JSON(400, gin.H{"error": err_1.Error()})
	}

	db := database.InitDB()

	userRepository := repositories.UserRepository{Db: db}

	user_1 := models.User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Balance: user.Balance,
	}
	err_2 := userRepository.CreateUser(&user_1)
	if err_2 != nil{
		panic(err_2)
	}
}