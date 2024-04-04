package postgres

import (
	"context"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
)

func (r *AccountsRepository) List(ctx context.Context) ([]dto.ResponseAccount, error) {
	var ac []dto.ResponseAccount
	r.Client.DB.Model(&entity.Account{}).Find(&ac)

	return ac, nil
}
