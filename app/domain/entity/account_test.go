package entity_test

import (
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/entity"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAccount_SetPassword(t *testing.T) {
	account := &entity.Account{}
	password := "newpassword123"

	account.SetPassword(password)

	assert.NotEqual(t, account.Secret, password)

	err := bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(password))
	assert.NoError(t, err)
}

func TestAccount_ValidatePassword(t *testing.T) {
	validPassword := "secret123"
	account := &entity.Account{}
	account.SetPassword(validPassword)

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{"valid password", validPassword, true},
		{"invalid password", "wrongpassword", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := account.ValidatePassword(tt.password)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAccount_IsValid(t *testing.T) {
	account := &entity.Account{
		Cpf: "12345678901",
	}

	err := account.IsValid()
	require.Error(t, err)
	assert.Equal(t, err, erring.ErrAccountCPFIsInvalid)
}

func TestAccount_Deposit(t *testing.T) {
	account := &entity.Account{Balance: 1000}

	tests := []struct {
		name    string
		amount  int
		wantErr bool
	}{
		{"positive deposit", 500, false},
		{"negative deposit", -100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account.Deposit(tt.amount)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, account.Balance, 1000+tt.amount)
			}
		})
	}
}

func TestAccount_Transfer(t *testing.T) {
	account1 := &entity.Account{Balance: 1000}
	account2 := &entity.Account{Balance: 500}

	tests := []struct {
		name             string
		amount           int
		wantErr          bool
		expectedBalance1 int
		expectedBalance2 int
	}{
		{"valid transfer", 200, false, 800, 700},
		{"insufficient balance", 2000, true, 1000, 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := account1.Transfer(tt.amount, account2)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, account1.Balance, tt.expectedBalance1)
				assert.Equal(t, account2.Balance, tt.expectedBalance2)
			}
		})
	}
}
