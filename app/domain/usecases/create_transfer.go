package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
	"github.com/LeonardoMuller13/digital-bank-api/app/pkg/uow"
)

type InputCreateTransfer struct {
	AccountDestinationCPF string
	Amount                int
}

func (u *UseCase) CreateTransfer(ctx context.Context, input InputCreateTransfer) error {
	if err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		repoAccount := u.getAccountRepository(ctx)
		repoTransfer := u.getTransferRepository(ctx)

		accountDest, err := repoAccount.GetAccountByCpf(ctx, input.AccountDestinationCPF)
		if err != nil {
			return erring.ErrTransferAccountDestinationNotFound
		}

		user, ok := ctx.Value(dto.UserKey).(*dto.User)
		if !ok {
			return errors.New("failed to assert context value as dto.User")
		}

		accountOrigin, err := repoAccount.GetAccountByID(ctx, int64(user.ID))
		if err != nil {
			return erring.ErrTransferUserNotFound
		}

		if err := accountOrigin.Transfer(input.Amount, &accountDest); err != nil {
			return fmt.Errorf("%w", err)
		}

		if err := repoTransfer.Create(ctx, int(accountOrigin.ID), int(accountDest.ID), input.Amount); err != nil {
			return fmt.Errorf("failed to create transfer from account %d to account %d with amount %d: %w", accountOrigin.ID, accountDest.ID, input.Amount, err)
		}

		if err := repoAccount.UpdateBalance(ctx, int(accountOrigin.ID), accountOrigin.Balance); err != nil {
			return fmt.Errorf("failed to update balance for account id %d: %w", accountOrigin.ID, err)
		}

		if err := repoAccount.UpdateBalance(ctx, int(accountDest.ID), accountDest.Balance); err != nil {
			return fmt.Errorf("failed to update balance for account id %d: %w", accountDest.ID, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
