package routers

import (
	"assignment-2/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.Home)
	router.POST("/order", handlers.CreateOrder)

	return router
}
