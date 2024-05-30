package main

import (
	"fmt"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/database"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	// database.MigrateDB(db)

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	// Use the custom CORS middleware
	r.Use(middleware.CORSMiddleware())
	fmt.Println("CORS okay")

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

	// common handlers(authenticated user)
	r.POST("/spin-roulette-v1/", handlers.SpinRouletteHandler)
	r.POST("/spin-roulette-v2/", handlers.UnfairSpinRouletteHandler)

	// Start the web server
	r.Run(":8081")
}
