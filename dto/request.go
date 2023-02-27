package dto

type Account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetAccount struct {
	ID string `json:"id"`
}

type DeleteAccount struct {
	ID string `json:"id"`
}
