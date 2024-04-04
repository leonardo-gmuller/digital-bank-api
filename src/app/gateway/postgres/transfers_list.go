package postgres

import (
	"context"
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
)

func (r *TransfersRepository) List(ctx context.Context) ([]entity.Transfer, error) {
	const (
		operation = "Repository.TransfersRepository.List"
	)
	var t []entity.Transfer
	result := r.Client.DB.Find(&t, "account_origin_id = ?", ctx.Value(dto.User{}).(dto.User).ID)
	if result.Error != nil {
		return []entity.Transfer{}, fmt.Errorf("%s -> %w", operation, result.Error)
	}

	return t, nil
}
