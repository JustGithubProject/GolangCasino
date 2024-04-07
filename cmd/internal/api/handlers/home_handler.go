package handlers

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/services"
	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
    services.HandleTemplate(c)
}