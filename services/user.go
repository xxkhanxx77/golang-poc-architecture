package services

import (
	"baac-backend/entity"
	"baac-backend/repo"
)

type UserService struct {
	repo *repo.Repo
}

func NewUserService(r *repo.Repo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) User() (*entity.User, error) {

	return s.repo.User()
}

func (s *UserService) CreateUser(user *entity.User) (*entity.User, error) {

	user, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
