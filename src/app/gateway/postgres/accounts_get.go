package postgres

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
)

func (r *AccountsRepository) GetAccountByID(ctx context.Context, id string) (entity.Account, error) {
	const (
		operation = "Repository.AccountsRepository.GetAccountByID"
	)
	var account entity.Account
	result := r.Client.DB.First(&account, id)

	if result.Error != nil {
		return entity.Account{}, fmt.Errorf("%s -> %w", operation, result.Error)
	}

	return account, nil
}

func (r *AccountsRepository) GetAccountByCpf(ctx context.Context, cpf string) (entity.Account, error) {
	const (
		operation = "Repository.AccountsRepository.GetAccountByID"
	)
	var account entity.Account
	result := r.Client.DB.First(&account, "cpf = ?", cpf)

	if result.Error != nil {
		return entity.Account{}, fmt.Errorf("%s -> %w", operation, result.Error)
	}

	return account, nil
}

func (r *AccountsRepository) GetAccountBalanceByID(ctx context.Context, id string) (dto.ResponseAccountBalance, error) {
	const (
		operation = "Repository.AccountsRepository.GetAccountBalanceByID"
	)
	var account entity.Account
	result := r.Client.DB.First(&account, id)

	if result.Error != nil {
		return dto.ResponseAccountBalance{}, fmt.Errorf("%s -> %w", operation, result.Error)
	}

	return dto.ResponseAccountBalance{
		ID:      int(account.ID),
		Name:    account.Name,
		Balance: account.Balance,
	}, nil
}
