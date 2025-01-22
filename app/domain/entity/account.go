package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
	"github.com/leonardo-gmuller/digital-bank-api/app/pkg/validations"
)

type Account struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"-"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (a *Account) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Secret), []byte(password))

	return err == nil
}

func (a *Account) IsValid() error {
	if !validations.CPFIsValid(a.Cpf) {
		return erring.ErrAccountCPFIsInvalid
	}

	return nil
}

func (a *Account) SetPassword(password string) {
	x := 8
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), x)
	a.Secret = string(hashedPassword)
}

func (a *Account) Deposit(amount int) error {
	if amount > 0 {
		a.Balance += amount

		return nil
	}

	return errors.New("deposit amount less than zero")
}

func (a *Account) Transfer(amount int, accountDestination *Account) error {
	if amount <= a.Balance {
		a.Balance -= amount

		if err := accountDestination.Deposit(amount); err != nil {
			return err
		}

		return nil
	}

	return erring.ErrTransferBalanceNotSufficient
}
