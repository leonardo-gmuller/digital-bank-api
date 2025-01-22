package repositories

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/entity"
)

func (r *AccountRepository) GetAccountByID(ctx context.Context, id int64) (entity.Account, error) {
	const (
		operation = "Repository.AccountRepository.GetAccountByID"
	)

	result, err := r.Queries.GetAccountById(ctx, id)
	if err != nil {
		return entity.Account{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return entity.Account{
		ID:        uint(result.ID),
		Name:      result.Name,
		Cpf:       result.Cpf,
		Secret:    result.Secret,
		Balance:   int(result.Balance),
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (r *AccountRepository) GetAccountByCpf(ctx context.Context, cpf string) (entity.Account, error) {
	const (
		operation = "Repository.AccountRepository.GetAccountByID"
	)

	result, err := r.Queries.GetAccountByCpf(ctx, cpf)
	if err != nil {
		return entity.Account{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return entity.Account{
		ID:        uint(result.ID),
		Name:      result.Name,
		Cpf:       result.Cpf,
		Secret:    result.Secret,
		Balance:   int(result.Balance),
		CreatedAt: result.CreatedAt.Time,
	}, nil
}

func (r *AccountRepository) GetAccountBalanceByID(ctx context.Context, id int64) (dto.ResponseAccountBalance, error) {
	const (
		operation = "Repository.AccountRepository.GetAccountBalanceByID"
	)

	result, err := r.Queries.GetAccountBalanceById(ctx, id)
	if err != nil {
		return dto.ResponseAccountBalance{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return dto.ResponseAccountBalance{
		ID:      int(id),
		Name:    result.Name,
		Balance: int(result.Balance),
	}, nil
}

func (r *AccountRepository) List(ctx context.Context) ([]dto.ResponseAccount, error) {
	const (
		operation = "Repository.TransfersRepository.List"
	)

	result, err := r.Queries.ListAccounts(ctx)
	if err != nil {
		return []dto.ResponseAccount{}, fmt.Errorf("%s -> %w", operation, err)
	}

	output := make([]dto.ResponseAccount, 0, len(result))

	for _, account := range result {
		output = append(output, dto.ResponseAccount{
			ID:   int(account.ID),
			Name: account.Name,
			Cpf:  account.Cpf,
		})
	}

	return output, nil
}
