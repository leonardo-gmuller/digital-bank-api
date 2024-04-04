package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
)

func (r *TransfersRepository) Create(ctx context.Context, input dto.RequestTransfer) error {
	const (
		operation = "Respository.TransfersRespository.Create"
	)
	tx := r.Client.DB.Begin()

	defer tx.Rollback()

	var accountOrigin entity.Account
	result := r.Client.DB.First(&accountOrigin, ctx.Value(dto.User{}).(dto.User).ID)
	if result.Error != nil {
		return fmt.Errorf("account origin not found")
	}

	var accountDestination entity.Account
	result = r.Client.DB.First(&accountDestination, "cpf = ?", input.Account_Destination_CPF)
	if result.Error != nil {
		return fmt.Errorf("account destination not found")
	}

	err := accountOrigin.Transfer(input.Amount, &accountDestination)
	if err != nil {
		log.Fatalf("%s -> %w", operation, err)
		return fmt.Errorf(err.Error())
	}

	tx.Save(&accountOrigin)
	tx.Save(&accountDestination)

	var newTransfer entity.Transfer
	newTransfer.AccountDestinationId = int(accountDestination.ID)
	newTransfer.AccountOriginId = int(accountOrigin.ID)
	newTransfer.Amount = input.Amount
	tx.Create(&newTransfer)
	tx.Commit()
	return nil
}
