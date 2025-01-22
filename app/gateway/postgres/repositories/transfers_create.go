package repositories

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/gateway/postgres/sqlc"
)

func (r *TransfersRepository) Create(ctx context.Context, accountOriginID int, accountDestinationID int, amount int) error {
	const (
		operation = "Repository.TransfersRespository.Create"
	)

	err := r.Queries.CreateTransfer(ctx, sqlc.CreateTransferParams{
		AccountOriginID:      int64(accountOriginID),
		AccountDestinationID: int64(accountDestinationID),
		Amount:               int32(amount),
	})
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}

	return nil
}
