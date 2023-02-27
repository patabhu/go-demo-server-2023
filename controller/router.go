package controller

import (
	"go-demo-server-2023/cmd/handler"
	"go-demo-server-2023/controller/api"
	"go-demo-server-2023/pkg/config"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandlers(cfg *config.Config, r *mux.Router, services *handler.Services) {
	res := api.New(services.UsecaseSvc)

	r.HandleFunc("/account", res.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/account/{account-id}", res.GetAccount).Methods(http.MethodGet)
	r.HandleFunc("/account/{account-id}", res.DeleteAccount).Methods(http.MethodDelete)
}
