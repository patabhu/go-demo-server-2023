package repository

import (
	"context"
	"database/sql"
	"go-demo-server-2023/model"
)

const (
	insertAccount = "INSERT INTO accounts (name, email) values (?, ?) "
	getAccount    = "SELECT id, name, email from accounts WHERE id = ? "
	deleteAccount = "DELETE FROM accounts where id = ? "
)

func (r *repository) SaveAccount(ctx context.Context, zr *model.Account) error {
	result, err := r.Conn.ExecContext(ctx, insertAccount, zr.Name, zr.Email)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *repository) GetAccount(ctx context.Context, id string) (*model.Account, error) {
	zr := model.Account{}
	if err := r.Conn.QueryRowContext(ctx, getAccount, id).Scan(
		&zr.ID,
		&zr.Name,
		&zr.Email,
	); err != nil {
		return nil, err
	}
	return &zr, nil
}

func (r *repository) DeleteAccount(ctx context.Context, id string) error {
	result, err := r.Conn.ExecContext(ctx, deleteAccount, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
