package usecase

import (
	"context"
	"go-demo-server-2023/dto"
	"go-demo-server-2023/repository"
)

type UsecaseInterface interface {
	CreateAccount(ctx context.Context, acc *dto.Account) error
	GetAccount(ctx context.Context, id string) (*dto.Account, error)
	DeleteAccount(ctx context.Context, id string) error
}

type usecase struct {
	repoSvc repository.RepositoryInterface
}

func New(repoSvc repository.RepositoryInterface) UsecaseInterface {
	return &usecase{
		repoSvc: repoSvc,
	}
}
