package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialized Router
	router := gin.Default()
	//Initialized Routes
	initializeRoutes(router)
	// Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
