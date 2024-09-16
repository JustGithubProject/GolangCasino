package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
)

func RegisterHandler(c *gin.Context) {
	// Handler to register user
	services.HandleUserRegister(c)
}


func LoginHandler(c *gin.Context) {
	// Handler to login user
	services.HandleUserLogin(c)
}


func LogoutHandler(c *gin.Context) {
	// Handler to logout user
    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}