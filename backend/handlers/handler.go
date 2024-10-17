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

	token, err := h.service.CreateUser(user)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return

	}
	response.Success(c, "User created successfully", token)

}

func (h *UserHandler) CheckEmailPassword(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON request to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		response.BadRequestError(c, "Invalid request")
		return
	}

	// Call the CheckPassword service
	user, err := h.service.CheckPassword(requestBody.Email, requestBody.Password)
	if err != nil {
		response.BadRequestError(c, err.Error())
		return
	}

	response.Success(c, "User logged in successfully", user)
}
