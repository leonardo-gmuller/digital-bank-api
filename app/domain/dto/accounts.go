package dto

type ResponseAccount struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

type ResponseAccountBalance struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
