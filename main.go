package main

import (
	"assignment-2/database"
	"assignment-2/routers"
)

func main() {
	PORT := ":8080"

	_, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	// migrate database
	database.Migrate()

	routers.SetupRouter().Run(PORT)

}
