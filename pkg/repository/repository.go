package repository

import (
	app "avito-dynamic-segment-back"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user app.User) (int, error)
}

type Segment interface {
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
