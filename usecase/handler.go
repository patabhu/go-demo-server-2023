package usecase

import (
	"context"
	"go-demo-server-2023/dto"
	"go-demo-server-2023/model"
)

func (u *usecase) CreateAccount(ctx context.Context, acc *dto.Account) error {
	return u.repoSvc.SaveAccount(ctx, &model.Account{
		Name:  acc.Name,
		Email: acc.Email,
	})
}

func (u *usecase) GetAccount(ctx context.Context, id string) (*dto.Account, error) {
	acc, err := u.repoSvc.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.Account{
		ID:    acc.ID,
		Name:  acc.Name,
		Email: acc.Email,
	}, nil
}

func (u *usecase) DeleteAccount(ctx context.Context, id string) error {
	return u.repoSvc.DeleteAccount(ctx, id)
}
