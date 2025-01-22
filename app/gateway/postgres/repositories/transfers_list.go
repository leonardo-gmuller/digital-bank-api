package repositories

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/dto"
)

func (r *TransfersRepository) List(ctx context.Context, id int) ([]dto.OutputTransfer, error) {
	const (
		operation = "Repository.TransfersRepository.List"
	)

	result, err := r.Queries.ListTransfer(ctx, int64(id))
	if err != nil {
		return []dto.OutputTransfer{}, fmt.Errorf("%s -> %w", operation, err)
	}

	output := make([]dto.OutputTransfer, 0, len(result))
	for _, transfer := range result {
		output = append(output, dto.OutputTransfer{
			AccountDestinationCPF: transfer.AccountDestinationCpf,
			Amount:                int(transfer.Amount),
			CreatedAt:             transfer.CreatedAt.Time,
		})
	}

	return output, nil
}
