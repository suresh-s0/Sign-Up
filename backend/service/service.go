package service

import (
	"backend/domain"
	"backend/models"
	"errors"
)

type Service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) domain.UserService {
	return &Service{repo}
}

func (s *Service) CreateUser(user models.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil

}

func (s *Service) CheckPassword(email string, password string) (models.User, error) {
	var user models.User

	// Assume you're fetching the user from DB and checking the password
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	// Add password verification logic here
	if user.Password != password {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}
