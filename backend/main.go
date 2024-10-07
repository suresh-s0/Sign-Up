package main

import (
	"backend/database"
	"fmt"
	"log"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to access the database object: %v", err)
	}
	defer sqlDB.Close()

	fmt.Println("Database connection successful!")
}
