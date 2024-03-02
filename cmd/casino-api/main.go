package main

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func main(){
	// Создание нового маршрутизатора Gin
	r := gin.Default()

	// Маршруты
	r.GET("/", handlers.HomeHandler)


	// Запуск веб-сервера
	r.Run(":8080")
}