package repository

import "github.com/jmoiron/sqlx"

type AppUser interface {
}

type AppSegment interface {
}

type Repository struct {
	AppUser
	AppSegment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
