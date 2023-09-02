package service

import (
	app "avito-dynamic-segment-back"
	"avito-dynamic-segment-back/pkg/repository"
)

type User interface {
	CreateUser(user app.User) (int, error)
}

type Segment interface {
}

type Service struct {
	User
	Segment
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repositories.User),
	}
}
