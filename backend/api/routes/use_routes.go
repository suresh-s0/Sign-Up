package routes

import (
	"backend/api/middleware"
	"backend/domain"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, repo domain.UserRepository, auth middleware.Middleware, ser domain.UserService) {
	handlers := handlers.NewHandler(ser)

	r.POST("/signup", handlers.CreateUser)
	r.POST("/signin", handlers.CheckEmailPassword)

}
