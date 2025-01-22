package dto

import "time"

type OutputTransfer struct {
	AccountDestinationCPF string    `json:"account_destination_cpf"`
	Amount                int       `json:"amount"`
	CreatedAt             time.Time `json:"created_at"`
}
