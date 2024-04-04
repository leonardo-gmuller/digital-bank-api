package dto

type RequestNewAccount struct {
	Name   string
	Cpf    string
	Secret string
}

type ResponseAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

type ResponseAccountBalance struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
