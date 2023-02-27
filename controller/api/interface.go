package api

import (
	"go-demo-server-2023/usecase"
	"net/http"
)

type ControllerInterface interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
}

func New(usecase usecase.UsecaseInterface) ControllerInterface {
	return &controller{
		usecase: usecase,
	}
}

type controller struct {
	usecase usecase.UsecaseInterface
}
