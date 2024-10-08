package domain

import "backend/models"

type UserRepository interface {
	CreateUser(user models.User) error
}

type UserService interface {
	CreateUser(user models.User) error
}
