package repositories

import (
	"context"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres"
	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres/sqlc"
)

type TransfersRepositoryInterface interface {
	Create(ctx context.Context, accountOriginID int, accountDestinationID int, amount int) error
	List(ctx context.Context, id int) ([]dto.OutputTransfer, error)
}

type TransfersRepository struct {
	*postgres.Client
	*sqlc.Queries
}

func NewTransfersRepository(client *postgres.Client) *TransfersRepository {
	return &TransfersRepository{
		client,
		sqlc.New(client.Pool),
	}
}
