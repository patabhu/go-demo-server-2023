package api

import (
	"encoding/json"
	"go-demo-server-2023/dto"
	"go-demo-server-2023/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *controller) CreateAccount(w http.ResponseWriter, r *http.Request) {

	req := &dto.Account{}
	resp := &dto.Response{
		Msg:   "failed to create account",
		Error: "failed to create account",
	}
	httpStatus := http.StatusBadRequest

	defer func() {
		if re := recover(); re != nil {
			err := re.(error)
			resp.Error = err.Error()
		}
		utils.JsonResponse(w, httpStatus, &resp)
	}()

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	err = c.usecase.CreateAccount(r.Context(), req)
	if err != nil {
		resp.Error = err.Error()
		return
	}
	resp.Error = ""
	resp.Msg = "account created"
	httpStatus = http.StatusOK
}

func (c *controller) GetAccount(w http.ResponseWriter, r *http.Request) {

	resp := &dto.GetAccountResponse{
		Msg:   "failed to get account",
		Error: "failed to get account",
	}

	httpStatus := http.StatusBadRequest

	defer func() {
		if re := recover(); re != nil {
			err := re.(error)
			resp.Error = err.Error()
		}
		utils.JsonResponse(w, httpStatus, &resp)
	}()

	id := mux.Vars(r)["account-id"]
	if id == "" {
		msg := "invalid account-id"
		resp.Error = msg
		resp.Msg = msg
		return
	}

	acc, err := c.usecase.GetAccount(r.Context(), id)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		resp.Error = err.Error()
		return
	}

	resp.Error = ""
	resp.Msg = "account info"
	resp.Data = acc
	httpStatus = http.StatusOK
}

func (c *controller) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	resp := &dto.Response{
		Msg:   "failed to delete account",
		Error: "failed to delete account",
	}
	httpStatus := http.StatusBadRequest

	defer func() {
		if re := recover(); re != nil {
			err := re.(error)
			resp.Error = err.Error()
		}
		utils.JsonResponse(w, httpStatus, &resp)
	}()

	id := mux.Vars(r)["account-id"]
	if id == "" {
		msg := "invalid account-id"
		resp.Error = msg
		resp.Msg = msg
		return
	}

	err := c.usecase.DeleteAccount(r.Context(), id)
	if err != nil {
		httpStatus = http.StatusInternalServerError
		resp.Error = err.Error()
		return
	}

	resp.Error = ""
	resp.Msg = "account deleted"
	httpStatus = http.StatusOK
}
