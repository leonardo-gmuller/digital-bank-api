package uow

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres"
)

type RepositoryFactory func(tx pgx.Tx) interface{}

type Interface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback(ctx context.Context) error
	Rollback(ctx context.Context) error
	Unregister(name string)
}

type Uow struct {
	DB           *postgres.Client
	Tx           pgx.Tx
	Repositories map[string]RepositoryFactory
}

func NewUow(db *postgres.Client) *Uow {
	return &Uow{
		DB:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}
