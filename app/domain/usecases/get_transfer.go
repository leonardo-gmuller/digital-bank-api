package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
)

type OutputListTransfer struct {
	Transfers []dto.OutputTransfer
}

func (u *UseCase) ListUserTransfer(ctx context.Context) (OutputListTransfer, error) {
	user, ok := ctx.Value(dto.UserKey).(*dto.User)
	if !ok {
		return OutputListTransfer{}, errors.New("failed to assert context value as dto.User")
	}

	repo := u.getTransferRepository(ctx)

	defer func() {
		_ = u.Uow.CommitOrRollback(ctx)
	}()

	result, err := repo.List(ctx, user.ID)
	if err != nil {
		return OutputListTransfer{}, fmt.Errorf("failed to list transfers for id %d: %w", user.ID, err)
	}

	return OutputListTransfer{result}, nil
}
