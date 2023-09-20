package main

import (
	"assignment-2/database"
	"assignment-2/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.InitDB()

	if err != nil {
		panic(err)
	}
	// fmt.Println("Sukses konek database")
	// db.AutoMigrate(&models.Order{}, &models.Item{})

	router := gin.Default()

	router.GET("/", handlers.Home)
	router.POST("/order", handlers.CreateOrder)

	fmt.Println("Listening on port 8080")
	router.Run(":8080")

}
