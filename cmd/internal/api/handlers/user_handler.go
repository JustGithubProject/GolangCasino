package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/repositories"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/models"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)




func CreateUserHandler(c *gin.Context) {
    // Parse the JSON data from the request body into the user model
    var user models.User
    err := c.BindJSON(&user)
    if err != nil {
        // If there's an error parsing JSON, return a 400 Bad Request response
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

    // Create a new user object with the parsed data
    user_1 := models.User{
        Name:     user.Name,
        Email:    user.Email,
        Password: user.Password,
        Balance:  user.Balance,
    }

    // Call the repository method to create the user in the database
    err = userRepository.CreateUser(&user_1)
    if err != nil {
        // If there's an error creating the user, panic
        panic(err)
    }
}



func GetUserByIdHandler(c *gin.Context) {
    // Extract the user ID from the request parameters
    userIDStr := c.Param("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        // If the user ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

    // Call the repository method to get the user by their ID
    user, err := userRepository.GetUserById(uint(userID))
    if err != nil {
        // If there's an error getting the user, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    // Return the user object as JSON
    c.JSON(http.StatusOK, user)
}


func UpdateUserHandler(c *gin.Context) {
    // Extract the user ID from the request parameters
    userIDStr := c.Param("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        // If the user ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

    // Call the repository method to get the user by their ID
    modelUser, err := userRepository.GetUserById(uint(userID))
    if err != nil {
        // If there's an error getting the user, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    // Call the repository method to update the user
    err = userRepository.UpdateUser(modelUser)
    if err != nil {
        // If there's an error updating the user, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    // Return a success message
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


func DeleteUserHandler(c *gin.Context) {
    // Extract the user ID from the request parameters
    userIDStr := c.Param("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        // If the user ID is not a valid integer, return a 400 Bad Request response
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Initialize the database connection
    db := database.InitDB()
    userRepository := repositories.UserRepository{Db: db}

    // Call the repository method to get the user by their ID
    modelUser, err := userRepository.GetUserById(uint(userID))
    if err != nil {
        // If there's an error getting the user, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
        return
    }

    // Call the repository method to delete the user
    err = userRepository.DeleteUser(modelUser)
    if err != nil {
        // If there's an error deleting the user, return a 500 Internal Server Error response
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    // Return a success message
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
