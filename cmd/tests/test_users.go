package tests

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/JustGithubProject/GolangCasino/cmd/internal/handlers"
    "github.com/JustGithubProject/GolangCasino/cmd/internal/models"
    "github.com/gin-gonic/gin"
)

func TestCreateUserHandler(t *testing.T) {
    // Создание тестового пользователя для отправки в запросе
    testUser := models.User{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "testpassword",
        Balance:  100,
    }

    // Инициализация тестового контекста Gin
    router := gin.Default()
    router.POST("/users", handlers.CreateUserHandler)

    // Создание тела JSON запроса из тестового пользователя
    requestBody := bytes.NewBuffer([]byte(`{
        "name": "Test User",
        "email": "test@example.com",
        "password": "testpassword",
        "balance": 100
    }`))

    // Создание HTTP POST запроса с JSON телом
    req, err := http.NewRequest("POST", "/users", requestBody)
    if err != nil {
        t.Fatalf("failed to create request: %v", err)
    }
    req.Header.Set("Content-Type", "application/json")

    // Запись ответа HTTP
    recorder := httptest.NewRecorder()

    // Обработка запроса с помощью тестового контекста Gin
    router.ServeHTTP(recorder, req)

    // Проверка статуса ответа
    if recorder.Code != http.StatusOK {
        t.Errorf("expected status %d; got %d", http.StatusOK, recorder.Code)
    }

 
}
