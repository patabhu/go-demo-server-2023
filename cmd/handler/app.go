package handler

import (
	"context"
	"fmt"
	"go-demo-server-2023/pkg/config"
	"go-demo-server-2023/pkg/mysql"
	"go-demo-server-2023/repository"
	"go-demo-server-2023/usecase"
)

type Services struct {
	UsecaseSvc usecase.UsecaseInterface
}

func New(ctx context.Context, cfg *config.Config) (*Services, error) {

	var (
		services Services
		err      error
	)

	mysqlConn, err := mysql.MysqlConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed mysql connection: %w", err)
	}
	repoSvc := repository.NewRepo(mysqlConn)

	services.UsecaseSvc = usecase.New(repoSvc)

	return &services, nil
}
