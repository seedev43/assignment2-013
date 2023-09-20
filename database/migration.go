package database

import (
	"assignment-2/models"
	"fmt"
)

func Migrate() {
	fmt.Println("Proses Migrate Database...")

	DB.AutoMigrate(&models.Order{}, &models.Item{})

	fmt.Println("Sukses Migrate Database...")
}
