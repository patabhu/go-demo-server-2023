package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

var (
	ErrInvalidFileExtension = errors.New("file extension not supported")
)

type Config struct {
	App struct {
		Env     string `yaml:"app_env" env:"APP_ENV"`
		Level   string `yaml:"app_level" env:"APP_LEVEL"`
		ApiPort string `yaml:"app_api_port" env:"APP_API_PORT"`
	}
	MySql struct {
		Username string `yaml:"mysql_username" env:"MYSQL_USERNAME"`
		Password string `yaml:"mysql_password" env:"MYSQL_PASSWORD"`
		Host     string `yaml:"mysql_host" env:"MYSQL_HOST"`
		Port     int    `yaml:"mysql_port" env:"MYSQL_PORT"`
		DBName   string `yaml:"mysql_dbname" env:"MYSQL_DBNAME"`
	}
}

func Load(fileNames ...string) (*Config, error) {

	loadFiles := make([]string, 0, len(fileNames))
	envFiles := make([]string, 0, len(fileNames))

	for _, file := range fileNames {
		fileParts := strings.Split(file, ".")
		fileExtn := fileParts[len(fileParts)-1]

		switch fileExtn {
		case "yml", "json", "yaml", "toml":
			loadFiles = append(loadFiles, file)
		case "env":
			envFiles = append(loadFiles, file)
		default:
			return nil, ErrInvalidFileExtension
		}
	}

	if len(envFiles) > 0 {
		err := godotenv.Load(envFiles...)
		if err != nil {
			return nil, fmt.Errorf("error while loading env files(%s): %w", strings.Join(envFiles, ","), err)
		}
	}

	cfg, err := loadConfig(loadFiles...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("loading %+v \n", cfg)
	return cfg, nil
}

func loadConfig(fileNames ...string) (*Config, error) {
	var config Config

	err := configor.Load(&config, fileNames...)
	if err != nil {
		return nil, fmt.Errorf("cannot load config files(%s): %w", strings.Join(fileNames, ","), err)
	}

	return &config, nil

}
