package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go-demo-server-2023/cmd/handler"
	"go-demo-server-2023/controller"
	"go-demo-server-2023/pkg/config"
	"go-demo-server-2023/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

const (
	defaultConfigFile = "./.env"
)

func showHelp() {
	flag.Usage()
	fmt.Println("Arguments: -service = api")
	fmt.Println("Arguments: -config = configfiles")
}

func main() {
	defer showHelp()

	configFile := flag.String("config", defaultConfigFile, "-config = [file] or -config = [files] (files seperated by ,)")
	service := flag.String("service", "", "buildName -service = [servicename]")
	flag.Parse()
	files := strings.Split(*configFile, ",")
	cfg, err := config.Load(files...)

	if err != nil {
		fmt.Println(err)
		return
	}
	logger.InitLogger(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch *service {
	case "api":
		buildHandler(cfg, ctx)

	default:
		zap.L().Error("service name not provided")
	}
}

func buildHandler(cfg *config.Config, ctx context.Context) {
	services, err := handler.New(ctx, cfg)
	if err != nil {
		zap.L().Fatal("failed to build api handler", zap.Error(err))
	}

	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	})
	controller.RegisterHandlers(cfg, router, services)

	handler := c.Handler(router)
	zap.L().Info("Handler service started", zap.String("port", cfg.App.ApiPort), zap.Any("time", time.Now()))
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.App.ApiPort), handler)
	if err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
