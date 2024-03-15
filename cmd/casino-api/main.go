package main

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"

)

func main(){
	db := database.InitDB()
	// database.MigrateDB(db)
	
	// Создание нового маршрутизатора Gin
	r := gin.Default()

	// Передача экземпляра базы данных в обработчики запросов
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// Маршруты
	r.GET("/", handlers.HomeHandler)
	r.POST("/create/user/", handlers.CreateUserHandler)
	r.GET("/user/:id", handlers.GetUserByIdHandler)
	r.PUT("/update/user", handlers.UpdateUserHandler)
	r.DELETE("/delete/user/:id", handlers.DeleteUserHandler)

	r.POST("/create/game/", handlers.CreateGameHandler)
	r.GET("/game/:id", handlers.GetGameByIdHandler)
	r.PUT("/update/game", handlers.UpdateGameHandler)
	r.DELETE("/delete/game/:id", handlers.DeleteGameHandler)

	// Запуск веб-сервера
	r.Run(":8080")
}