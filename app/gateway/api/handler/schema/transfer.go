package schema

type RequestCreateTransfer struct {
	AccountDestinationCPF string `json:"account_destination_cpf"`
	Amount                int    `json:"amount"`
}
