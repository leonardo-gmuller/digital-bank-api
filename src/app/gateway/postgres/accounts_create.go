package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/helpers"
)

func (r *AccountsRepository) Create(ctx context.Context, input dto.RequestNewAccount) error {
	const (
		operation = "Repository.AccountsRepository.Create"
	)
	var account entity.Account
	if !helpers.CPFIsValid(input.Cpf) {
		return fmt.Errorf("CPF is not valid")
	}
	result := r.Client.DB.First(&account, "cpf = ?", input.Cpf)
	if result.Error == nil {
		log.Fatalf("%s -> %w", operation, result.Error)
		return fmt.Errorf("account already exists")
	}
	account.Name = input.Name
	account.Cpf = input.Cpf
	account.Balance = 1000 //For tests, init account with 1000
	account.SetPassword(input.Secret)
	r.Client.DB.Create(&account)

	return nil
}
