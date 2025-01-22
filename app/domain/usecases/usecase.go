package usecases

import (
	"context"

	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/postgres/repositories"
	"github.com/leonardo-gmuller/digital-bank-api/app/pkg/uow"
)

type UseCase struct {
	AppName string

	// UOW
	Uow uow.Interface
}

func (u *UseCase) getAccountRepository(ctx context.Context) repositories.AccountRepositoryInterface {
	repo, err := u.Uow.GetRepository(ctx, "AccountRepository")
	if err != nil {
		panic(err)
	}

	repoAsAccountRepo, ok := repo.(repositories.AccountRepositoryInterface)
	if !ok {
		return nil
	}

	return repoAsAccountRepo
}

func (u *UseCase) getTransferRepository(ctx context.Context) repositories.TransfersRepositoryInterface {
	repo, err := u.Uow.GetRepository(ctx, "TransferRepository")
	if err != nil {
		panic(err)
	}

	repoAsTransferRepo, ok := repo.(repositories.TransfersRepositoryInterface)
	if !ok {
		return nil
	}

	return repoAsTransferRepo
}
