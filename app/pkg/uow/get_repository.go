package uow

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.DB.Pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to begin transaction: %w", err)
		}

		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx)

	return repo, nil
}
