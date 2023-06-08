package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma nova instância do gin
	router := gin.Default()

	// Rota GET "/hello"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá, Mundo!",
		})
	})

	// Inicia o servidor HTTP
	router.Run(":8080")
}
