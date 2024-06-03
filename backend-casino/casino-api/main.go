package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/database"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

func LogResponseHeaders() gin.HandlerFunc {
	fmt.Println("LogResponseHeader before return func")
	return func(c *gin.Context) {
		// Попытка логирования заголовков до c.Next()
		fmt.Println("До c.Next()")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				fmt.Printf("До c.Next() - %s: %s\n", key, value)
			}
		}

		c.Next() // Обработка запроса

		// Логирование заголовков после c.Next()
		fmt.Println("После c.Next()")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				fmt.Printf("После c.Next() - %s: %s\n", key, value)
			}
		}
	}
}


func main() {
	db := database.InitDB()
	// database.MigrateDB(db)

	r := gin.Default()

	// Использование cors middleware из библиотеки gin-contrib/cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	fmt.Println("CORS middleware applied")

	r.Use(middleware.LoggerMiddleware())

	// Добавление middleware для логирования заголовков
	r.Use(LogResponseHeaders())
	fmt.Println("LogResponseHeaders is empty?")
	// Passing the database instance to query handlers
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// Routes for management(admin)
	r.GET("/user/:id", handlers.GetUserByIdHandler)
	r.GET("/user/name/:username", handlers.GetUserByUsernameHandler)
	r.PUT("/update/user", handlers.UpdateUserHandler)
	r.DELETE("/delete/user/:id", handlers.DeleteUserHandler)
	r.POST("/create/game/", handlers.CreateGameHandler)
	r.GET("/game/:id", handlers.GetGameByIdHandler)
	r.PUT("/update/game", handlers.UpdateGameHandler)
	r.DELETE("/delete/game/:id", handlers.DeleteGameHandler)

	// auth_handlers
	r.POST("/register/user/", handlers.RegisterHandler)
	r.POST("/login/user/", handlers.LoginHandler)
	r.POST("/logout", handlers.LogoutHandler)

	// common handlers(authenticated user)
	r.POST("/spin-roulette-v1/", handlers.SpinRouletteHandler)
	r.POST("/spin-roulette-v2/", handlers.UnfairSpinRouletteHandler)

	// Start the web server
	r.Run(":8081")
}
