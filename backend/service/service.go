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

func (s *Service) CreateUser(user models.User) error {
	err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	// token, _ := middleware.CreateJWt(uint(user.Id))
	// return token, nil
	return nil

}

func (s *Service) CheckPassword(email string, password string) (*models.User, error) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}
	if user == nil {
		return &models.User{}, errors.New("user not found")
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	token, err := middleware.CreateJWt(uint(user.Id))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	if err = s.repo.UpdateToken(email, token); err != nil {
		return &models.User{}, err
	}

	user.Token = token

	return user, nil
}
