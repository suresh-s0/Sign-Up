package repository

import (
	"backend/domain"
	"backend/models"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) domain.UserRepository {

	return &UserRepository{db}
}

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) CreateUser(user models.User) error {

	err := r.db.Exec("INSERT INTO users(name, email, password, is_admin) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Password, user.IsAdmin).Error

	if err != nil {
		return err
	}

	return nil

}
