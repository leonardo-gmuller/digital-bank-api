package tests

import (
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	account := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Secret:  "123456",
		Balance: 1000,
	}
	account.Deposit(1000)
	assert.Equal(t, 2000, account.Balance)
}

func TestTransfer(t *testing.T) {
	accountOrigin := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Secret:  "123456",
		Balance: 1000,
	}
	accountDestination := models.Account{
		Name:    "Teste",
		Cpf:     "12345678901",
		Secret:  "123456",
		Balance: 1000,
	}
	accountOrigin.Transfer(100, &accountDestination)
	assert.Equal(t, 900, accountOrigin.Balance)
	assert.Equal(t, 1100, accountDestination.Balance)
}
