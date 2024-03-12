package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)



func CreateUserHandler(
	c *gin.Context,
){
	var user models.User
	
	// Этот код выполняет привязку JSON данных из тела HTTP запроса к структуре user.
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


func GetUserByIdHandler(c *gin.Context) {
    userIDStr := c.Param("id")

    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    db := database.InitDB()

    userRepository := repositories.UserRepository{Db: db}
    user, err := userRepository.GetUserById(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
    userIDStr := c.Param("id")

    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

    modelUser, err := userRepository.GetUserById(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    // Обновляем данные пользователя
    err = userRepository.UpdateUser(modelUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

