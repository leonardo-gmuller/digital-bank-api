package app

import (
	"github.com/LeonardoMuller13/digital-bank-api/src/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/gateway/postgres"
)

type App struct {
	UseCase *usecases.UseCase
}

func New(config config.Config, db *postgres.Client) (*App, error) { //nolint: revive
	useCase := &usecases.UseCase{
		AppName:             config.App.Name,
		AccountsRepository:  postgres.NewAccountsRepository(db),
		TransfersRepository: postgres.NewTransfersRepository(db),
	}

	return &App{
		UseCase: useCase,
	}, nil
}
