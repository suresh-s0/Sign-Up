package service

import (
	"backend/api/middleware"
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

func (s *Service) CreateUser(user models.User) (string, error) {
	err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}
	token, _ := middleware.CreateJWt(uint(user.Id))
	return token, nil

}

func (s *Service) CheckPassword(email string, password string) (*models.User, error) {
	var user *models.User

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return &models.User{}, err
	}

	if user.Password != password {
		return &models.User{}, errors.New("invalid password")
	}

	return user, nil
}
