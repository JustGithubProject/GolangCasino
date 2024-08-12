package main

import (
	"time"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers/paypal"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/database"
)

func LogResponseHeaders() gin.HandlerFunc {
	log.Println("LogResponseHeader before return func")
	return func(c *gin.Context) {
		log.Println("До c.Next()")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				log.Printf("До c.Next() - %s: %s\n", key, value)
			}
		}

		c.Next()

		// Logging headers after c.Next()
		log.Println("После c.Next()")
		for key, values := range c.Writer.Header() {
			for _, value := range values {
				log.Printf("После c.Next() - %s: %s\n", key, value)
			}
		}
	}
}


func main() {
	db := database.InitDB()
	sqlDB, err := db.DB()
    if err != nil{
        log.Fatal("Failed to get database connection", err)
    }

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

    defer sqlDB.Close()
	// database.MigrateDB(db)

	r := gin.Default()

	// Using cors middleware from the gin-contrib/cors library
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	log.Println("CORS middleware applied")

	// Adding middleware for logging headers
	r.Use(LogResponseHeaders())
	log.Println("LogResponseHeaders is empty?")

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
	r.POST("/spin-roulette-v2/", handlers.UnfairSpinRouletteHandler) // like rooms(6 different handlers)
	r.POST("/spin-roulette-v3/", handlers.VeryBadSpinRouletteHandler)
	r.POST("/spin-slot-v1/", handlers.SweetBonanzaSlotHandler)


	// paypal handlers
	r.POST("/paypal/create/order/", paypal_handlers.CreatePaymentOrder) // Correct way to add balance for users
	r.GET("/paypal/info/order/", paypal_handlers.GetOrderDetailByID) 
	
	// Start the web server
	r.Run(":8081")
}
