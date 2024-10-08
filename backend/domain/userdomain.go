package domain

import "backend/models"

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type UserService interface {
	CreateUser(user models.User) error
	CheckPassword(email string, password string) (models.User, error)
}
