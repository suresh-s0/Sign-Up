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

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	// Use a pointer in Scan to populate the user object
	err := r.db.Raw("SELECT name, email, password, is_admin FROM users WHERE email = ?", email).Scan(&user).Error
	if err != nil {
		return models.User{}, err // Return an empty User struct and the error
	}
	return user, nil
}
