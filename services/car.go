package services

import (
	"baac-backend/entity"
	"baac-backend/repo"
)

type CarService struct {
	repo *repo.Repo
}

func NewCarService(r *repo.Repo) *CarService {
	return &CarService{repo: r}
}

func (s *CarService) Car() (*entity.Car, error) {

	return s.repo.Car()
}
