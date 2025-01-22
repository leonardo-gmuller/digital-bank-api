package uow

import (
	"context"
	"errors"
)

func (u *Uow) Rollback(ctx context.Context) error {
	if u.Tx == nil {
		return errors.New("no transaction to rollback")
	}

	err := u.Tx.Rollback(ctx)
	if err != nil {
		return errors.New(err.Error())
	}

	u.Tx = nil

	return nil
}
