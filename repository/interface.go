package repository

import (
	"context"
	"database/sql"

	"go-demo-server-2023/model"
)

type RepositoryInterface interface {
	SaveAccount(ctx context.Context, zr *model.Account) error
	GetAccount(ctx context.Context, id string) (*model.Account, error)
	DeleteAccount(ctx context.Context, id string) error
}

type repository struct {
	Conn *sql.DB
}

func NewRepo(conn *sql.DB) RepositoryInterface {
	return &repository{
		Conn: conn,
	}
}
