package uow

import (
	"context"
	"fmt"
)

func (u *Uow) CommitOrRollback(ctx context.Context) error {
	err := u.Tx.Commit(ctx)
	if err != nil {
		if errRb := u.Rollback(ctx); errRb != nil {
			return fmt.Errorf("original error: %w, rollback error: %w", err, errRb)
		}

		return fmt.Errorf("commit error: %w", err)
	}

	u.Tx = nil

	return nil
}
