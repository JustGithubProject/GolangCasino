package handlers

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/services"
	"github.com/gin-gonic/gin"
)




func RegisterHandler(c *gin.Context) {
	services.HandleUserRegister(c)
}


func LoginHandler(c *gin.Context) {
	services.HandleUserLogin(c)
}