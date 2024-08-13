package database

import (
    "log"
    "os"

    "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/JustGithubProject/GolangCasino/backend-casino/internal/models"
)


func MigrateDB(db *gorm.DB){
    db.AutoMigrate(&models.User{}, &models.Game{}, &models.Payment{})
}

var DB *gorm.DB

func InitDB() *gorm.DB {
    
    // Loading .env file
    err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env file")
    }

    // Getting variable of .env 
    dsn := os.Getenv("DB_DSN")


    // Setting of connection
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: dsn,
    }), &gorm.Config{})

    if err != nil {
        log.Fatal("failed to connect database")
    }


    DB = db
    return db
}