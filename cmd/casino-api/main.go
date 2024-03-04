package main

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
)

func main(){
	db := database.InitDB()
	
	// Создание нового маршрутизатора Gin
	r := gin.Default()

	// Передача экземпляра базы данных в обработчики запросов
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// Маршруты
	r.GET("/", handlers.HomeHandler)


	// Запуск веб-сервера
	r.Run(":8080")
}