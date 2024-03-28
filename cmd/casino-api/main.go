package main

import (
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/database"
	"github.com/JustGithubProject/GolangCasino/cmd/internal/api/middleware"

)

func main(){
	db := database.InitDB()
	// database.MigrateDB(db)
	
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	// Passing the database instance to query handlers
	r.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// Routes
	r.GET("/", handlers.HomeHandler)
	r.POST("/create/user/", handlers.CreateUserHandler)
	r.GET("/user/:id", handlers.GetUserByIdHandler)
	r.PUT("/update/user", handlers.UpdateUserHandler)
	r.DELETE("/delete/user/:id", handlers.DeleteUserHandler)

	r.POST("/create/game/", handlers.CreateGameHandler)
	r.GET("/game/:id", handlers.GetGameByIdHandler)
	r.PUT("/update/game", handlers.UpdateGameHandler)
	r.DELETE("/delete/game/:id", handlers.DeleteGameHandler)


	// auth_handlers
	r.POST("/register/user/", handlers.RegisterHandler)
	r.POST("/login/user/",  handlers.LoginHandler)

	// Start the web server
	r.Run(":8080")
}