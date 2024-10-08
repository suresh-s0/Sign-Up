package handlers

import (
	"backend/api/response"
	"backend/domain"
	"backend/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domain.UserService
}

func NewHandler(service domain.UserService) *UserHandler {
	return &UserHandler{service}

}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		response.BadRequestError(c, err.Error())
		return

	}
	response.Success(c, "User created successfully", user)

}
