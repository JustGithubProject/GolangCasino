package handlers

import (
	"github.com/gin-gonic/gin"
)


// HomeHandler обрабатывает запросы к главной странице
func HomeHandler(c *gin.Context){
	c.String(200, "Привет, мир!")
}	