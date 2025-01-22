package uow

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil {
		return errors.New("transaction already started")
	}

	tx, err := u.DB.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("error on begin transaction: %w", err)
	}

	u.Tx = tx

	err = fn(u)
	if err != nil {
		if errRb := u.Rollback(ctx); errRb != nil {
			return fmt.Errorf("original error: %w, rollback error: %w", err, errRb)
		}

		return err
	}

	return u.CommitOrRollback(ctx)
}
