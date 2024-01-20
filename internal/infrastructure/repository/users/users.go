package users

import (
	"api/internal/core"
	"context"

	"github.com/jmoiron/sqlx"
)

type users struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *users {
	return &users{db: db}
}

func (r *users) GetOne(ctx context.Context, ID int) (*core.User, error) {
	var output core.User
	row := r.db.QueryRowx("SELECT * FROM users WHERE id = $1", ID)
	if err := row.StructScan(&output); err != nil {
		return nil, err
	}
	return &output, nil
}
