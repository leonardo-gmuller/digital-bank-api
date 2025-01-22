package schema

type RequestNewAccount struct {
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}
