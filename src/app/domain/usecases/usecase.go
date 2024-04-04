package usecases

import (
	"context"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
)

type UseCase struct {
	AppName string

	// Repos
	AccountsRepository  accountsRepository
	TransfersRepository transfersRepository
}

type transfersRepository interface {
	Create(ctx context.Context, input dto.RequestTransfer) error
	List(ctx context.Context) ([]entity.Transfer, error)
}

type accountsRepository interface {
	GetAccountBalanceByID(ctx context.Context, id string) (dto.ResponseAccountBalance, error)
	GetAccountByID(ctx context.Context, id string) (entity.Account, error)
	GetAccountByCpf(ctx context.Context, cpf string) (entity.Account, error)
	Create(ctx context.Context, input dto.RequestNewAccount) error
	List(ctx context.Context) ([]dto.ResponseAccount, error)
}
