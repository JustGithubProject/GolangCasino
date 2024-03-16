package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)


func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Printf("Started handling request %s %s", c.Request.Method, c.Request.URL.Path)
        c.Next()
        log.Printf("Finished handling request")
    }
}