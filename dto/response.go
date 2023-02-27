package dto

type Response struct {
	Msg   string `json:"msg"`
	Error string `json:"error,omitempty"`
}

type GetAccountResponse struct {
	Msg   string   `json:"msg"`
	Error string   `json:"error,omitempty"`
	Data  *Account `json:"data,omitempty"`
}
