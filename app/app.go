package app

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/LeonardoMuller13/digital-bank-api/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres/repositories"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres/sqlc"
	"github.com/LeonardoMuller13/digital-bank-api/app/pkg/uow"
)

type App struct {
	UseCase *usecases.UseCase
}

func New(_ context.Context, config config.Config, db *postgres.Client) (*App, error) {
	uow := uow.NewUow(db)
	uow.Register("AccountRepository", func(tx pgx.Tx) interface{} {
		repo := repositories.NewAccountRepository(db)
		repo.Queries = sqlc.New(tx)

		return repo
	})
	uow.Register("TransferRepository", func(tx pgx.Tx) interface{} {
		repo := repositories.NewTransfersRepository(db)
		repo.Queries = sqlc.New(tx)

		return repo
	})

	useCase := &usecases.UseCase{
		AppName: config.App.Name,
		Uow:     uow,
	}

	return &App{
		UseCase: useCase,
	}, nil
}
