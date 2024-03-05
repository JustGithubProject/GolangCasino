package database

import (
    "log"
    "os"

    "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    // Загрузка переменных окружения из файла .env
    err := godotenv.Load()
    if err != nil{
        log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("DB_DSN")


    // Настройка соединения с базой данных PostgreSQL
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: dsn,
    }), &gorm.Config{})
    
    if err != nil {
        log.Fatal("failed to connect database")
    }
    
    DB = db
    return db
}