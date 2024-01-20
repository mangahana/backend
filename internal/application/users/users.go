package users

import "api/internal/infrastructure/repository"

type users struct {
	repo repository.Users
}

func New(repo repository.Users) *users {
	return &users{repo: repo}
}
