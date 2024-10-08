package routes

import (
	"backend/repository"
	"backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	apiRouter := r.Group("/api")
	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)

	UserRoutes(apiRouter, userRepo, userService)

}
