package domain

import "backend/models"

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
}

type UserService interface {
	CreateUser(user models.User) (string, error)
	CheckPassword(email string, password string) (*models.User, error)
}
