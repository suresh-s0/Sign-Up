package routes

import (
	"backend/domain"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, repo domain.UserRepository, ser domain.UserService) {
	handlers := handlers.NewHandler(ser)
	r.POST("/signup", handlers.CreateUser)
	r.POST("/signin", handlers.CheckEmailPassword)

}
