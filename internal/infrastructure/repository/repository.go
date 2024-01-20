package repository

import (
	"api/internal/infrastructure/repository/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users Users
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		Users: users.New(db),
	}
}

type Users interface {
	// GetOne(context.Context,)
}
