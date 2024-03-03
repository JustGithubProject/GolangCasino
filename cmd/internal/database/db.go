package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    // Настройка соединения с базой данных PostgreSQL
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
    }), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }
    
    DB = db
    return db
}