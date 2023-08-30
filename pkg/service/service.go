package service

import "avito-dynamic-segment-back/pkg/repository"

type AppUser interface {
}

type AppSegment interface {
}

type Service struct {
	AppUser
	AppSegment
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{}
}
