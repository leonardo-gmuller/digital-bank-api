package repositories

import (
	"context"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/entity"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/postgres"
	"github.com/leonardo-gmuller/digital-bank-api/app/gateway/postgres/sqlc"
)

type AccountRepositoryInterface interface {
	GetAccountBalanceByID(ctx context.Context, id int64) (dto.ResponseAccountBalance, error)
	GetAccountByID(ctx context.Context, id int64) (entity.Account, error)
	GetAccountByCpf(ctx context.Context, cpf string) (entity.Account, error)
	Create(ctx context.Context, name string, cpf string, secret string, balance int32) error
	UpdateBalance(ctx context.Context, id int, newBalance int) error
	List(ctx context.Context) ([]dto.ResponseAccount, error)
}

type AccountRepository struct {
	*postgres.Client
	*sqlc.Queries
}

func NewAccountRepository(client *postgres.Client) *AccountRepository {
	return &AccountRepository{client, sqlc.New(client.Pool)}
}
