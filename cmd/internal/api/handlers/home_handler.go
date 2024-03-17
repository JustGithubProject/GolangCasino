package handlers

import (
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)


func HomeHandler(c *gin.Context) {
    // Loading the contents of the HTML file
    tmpl, err := template.ParseFiles("D:/Users/Kropi/Desktop/All directory/go/casino/cmd/templates/index.html")
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

    // Send HTML page in response
    err = tmpl.Execute(c.Writer, nil)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
}