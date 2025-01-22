package repositories

import (
	"context"
	"fmt"

	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/postgres/sqlc"
)

func (r *AccountRepository) Create(ctx context.Context, name string, cpf string, secret string, balance int32) error {
	const (
		operation = "Repository.AccountRepository.Create"
	)

	err := r.Queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Name:    name,
		Cpf:     cpf,
		Secret:  secret,
		Balance: balance,
	})
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	return nil
}
