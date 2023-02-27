package mysql

import (
	"database/sql"
	"fmt"
	"go-demo-server-2023/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func MysqlConnection(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		cfg.MySql.Username,
		cfg.MySql.Password,
		cfg.MySql.Host,
		cfg.MySql.Port,
		cfg.MySql.DBName,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	zap.L().Info("MySQL Successfully Connected.")
	return db, nil
}
