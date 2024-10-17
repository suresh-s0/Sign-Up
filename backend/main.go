package main

import (
	"backend/api/routes"
	"backend/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	router := gin.Default()
	setupCors(router)
	setupRoutes(router, db)
	router.Run(":8080")

}

func setupRoutes(router *gin.Engine, db *gorm.DB) {
	routes.SetupRoutes(router, db)
}

func setupCors(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	router.Use(cors.New(corsConfig))
}
