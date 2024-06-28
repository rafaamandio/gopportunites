package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o Router utilizando as configurações Default do Gin
	router := gin.Default()
	//Definindo a rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos Rodando a nossa API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
