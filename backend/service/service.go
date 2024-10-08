package service

import (
	"backend/domain"
	"backend/models"
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
