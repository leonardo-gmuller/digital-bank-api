package repositories

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres/sqlc"
)

func (r *AccountRepository) UpdateBalance(ctx context.Context, id int, newBalance int) error {
	const (
		operation = "Repository.AccountRepository.UpdateBalance"
	)

	err := r.Queries.UpdateBalance(ctx, sqlc.UpdateBalanceParams{
		Balance: int32(newBalance),
		ID:      int64(id),
	})
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	return nil
}
