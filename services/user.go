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
