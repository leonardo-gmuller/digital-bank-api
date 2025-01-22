package usecases

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/entity"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
	"github.com/LeonardoMuller13/digital-bank-api/app/pkg/uow"
)

type InputCreateAccount struct {
	Name    string
	Cpf     string
	Secret  string
	Balance int
}

func (u *UseCase) CreateAccount(ctx context.Context, input InputCreateAccount) error {
	account := entity.Account{
		Name: input.Name,
		Cpf:  input.Cpf,
	}
	account.SetPassword(input.Secret)

	if err := account.IsValid(); err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := account.Deposit(input.Balance); err != nil {
		return fmt.Errorf("failed to deposit amount: %w", err)
	}

	if err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		repo := u.getAccountRepository(ctx)

		if _, err := repo.GetAccountByCpf(ctx, account.Cpf); err == nil {
			return erring.ErrAccountExists
		}

		err := repo.Create(ctx, account.Name, account.Cpf, account.Secret, int32(account.Balance))
		if err != nil {
			return fmt.Errorf("failed to create account on database: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
