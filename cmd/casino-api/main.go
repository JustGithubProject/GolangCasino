package main

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/middleware"
	"github.com/gin-gonic/gin"
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

	// CORS middleware with specific origins
	r.Use(func(c *gin.Context) {
		allowedOrigins := []string{"http://localhost:5173", "http://127.0.0.1:5173"}
		origin := c.Request.Header.Get("Origin")
		isAllowed := false
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Routes

	// Routes for management(admin)
	r.GET("/user/:id", handlers.GetUserByIdHandler)
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
