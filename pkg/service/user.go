package service

import (
	app "avito-dynamic-segment-back"
	"avito-dynamic-segment-back/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user app.User) (int, error) {
	return s.repo.CreateUser(user)
}
