package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


func HashPassword(password string) string {
    hasher := sha256.New()
    hasher.Write([]byte(password))
    hash := hasher.Sum(nil)
    return hex.EncodeToString(hash)
}

func CheckPasswordHash(password, hash string) bool {
    inputHash := HashPassword(password) // Здесь хешируем только исходный пароль
    return inputHash == hash
}

func main(){
    hashed_password := HashPassword("12345678")
    result := CheckPasswordHash("12345678", hashed_password) // Передаем исходный пароль, а не его хеш
    fmt.Println(result) // Выводится true
    fmt.Println(hashed_password == HashPassword("12345678")) // Также выводится true
}