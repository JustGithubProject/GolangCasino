package main

import (
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/database"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	db := database.InitDB()
	// database.MigrateDB(db)

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	// Passing the database instance to query handlers
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes

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
