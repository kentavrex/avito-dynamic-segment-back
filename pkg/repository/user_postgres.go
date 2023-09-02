package repository

import (
	app "avito-dynamic-segment-back"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user app.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) values($1) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
