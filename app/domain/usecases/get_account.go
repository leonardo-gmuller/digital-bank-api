package usecases

import (
	"context"
	"fmt"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/entity"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
)

type OutputListAccount struct {
	Accounts []dto.ResponseAccount `json:"accounts"`
}

func (u *UseCase) ListAccount(ctx context.Context) (OutputListAccount, error) {
	repo := u.getAccountRepository(ctx)

	defer func() {
		_ = u.Uow.CommitOrRollback(ctx)
	}()

	result, err := repo.List(ctx)
	if err != nil {
		return OutputListAccount{}, fmt.Errorf("failed to list accounts: %w", err)
	}

	return OutputListAccount{result}, nil
}

func (u *UseCase) GetBalanceByID(ctx context.Context, id int) (dto.ResponseAccountBalance, error) {
	repo := u.getAccountRepository(ctx)

	defer func() {
		_ = u.Uow.CommitOrRollback(ctx)
	}()

	result, err := repo.GetAccountBalanceByID(ctx, int64(id))
	if err != nil {
		return dto.ResponseAccountBalance{}, fmt.Errorf("failed to get account balance for account ID %d: %w", id, err)
	}

	return result, nil
}

func (u *UseCase) GetAccountByCpf(ctx context.Context, cpf string) (entity.Account, error) {
	repo := u.getAccountRepository(ctx)

	defer func() {
		_ = u.Uow.CommitOrRollback(ctx)
	}()

	result, err := repo.GetAccountByCpf(ctx, cpf)
	if err != nil {
		return entity.Account{}, erring.ErrAccountNotExists
	}

	return result, nil
}
