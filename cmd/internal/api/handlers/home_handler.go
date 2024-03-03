package handlers

import (
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
    // Загружаем содержимое HTML файла
    tmpl, err := template.ParseFiles("D:/Users/Kropi/Desktop/All directory/go/casino/cmd/templates/index.html")
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

    // Отправляем HTML страницу в ответе
    err = tmpl.Execute(c.Writer, nil)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
}