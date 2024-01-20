package application

import (
	"api/internal/application/users"
	"api/internal/infrastructure/repository"
	"context"
)

func New(repo repository.Users) *UseCase {
	return &UseCase{
		Users: users.New(repo),
	}
}

type UseCase struct {
	Users Users
}

type Users interface {
	SignUp(context.Context)
}
