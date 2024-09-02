package main

import (
	"time"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers/paypal"
	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/api/handlers/google-auth"
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
	// database.MigrateDB(db)
    defer sqlDB.Close()
	

	r := gin.Default()

	// Using cors middleware from the gin-contrib/cors library
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5173", "http://localhost:5173"},
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
	// slots handlers
	r.POST("/spin-slot-v1/", handlers.SweetBonanzaSlotHandler)
	r.POST("/spin-slot-v2/", handlers.DogHouseSlotHandler)
	r.POST("/spin-slot-v3/", handlers.WolfGoldSlotHandler)
	r.POST("/spin-slot-v4/", handlers.BigBassBonanzaSlotHandler)
	r.POST("/spin-slot-v5/", handlers.DiamondStrikeSlotHandler)

	// paypal handlers
	r.POST("/paypal/create/order/", paypal_handlers.CreatePaymentOrder) 
	r.GET("/paypal/info/order/", paypal_handlers.GetOrderDetailByID) 
	r.GET("/paypal/payments/history/", paypal_handlers.GetListPaypalPayments)
	r.POST("/paypal/update/approved/order/", paypal_handlers.UpdatePaymentStatusToApproved)
	r.POST("/paypal/update/pickup/money", paypal_handlers.PickUpMoneyAndStatusToSuccess)
	r.POST("/paypal/withdraw/funds/", paypal_handlers.WithdrawFundsPaypal)


	// Google auth handlers
	r.GET("/google/oauth/", google_auth.HandleGoogleLogin)
    r.POST("/google/auth/callback/", google_auth.HandleGoogleCallback)
	
	// Start the web server
	r.Run(":8081")
}
 