package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `gorm:"primaryKey";json:"id"`
	Name      string         `json:"name"`
	Cpf       string         `json:"cpf"`
	Secret    string         `json:"secret"`
	Balance   int            `json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index";json:"deleted_at"`
}

func (a *Account) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	a.Secret = string(hashedPassword)
}
func (a *Account) Deposit(amount int) (string, int) {
	if amount > 0 {
		a.Balance += amount
		return "Deposit made successfully!", int(a.Balance)
	}
	return "Deposit amount less than zero.", int(a.Balance)
}
func (a *Account) Transfer(amount int, accountDestination *Account) error {
	if amount <= a.Balance {
		a.Balance -= amount
		accountDestination.Deposit(amount)
		return nil
	}
	return fmt.Errorf("The originating account does not have sufficient balance.")
}

var Accounts []Account
