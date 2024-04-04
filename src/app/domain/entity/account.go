package entity

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" validate:"nonzero"`
	Cpf       string         `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	Secret    string         `json:"-"`
	Balance   int            `json:"balance, omitempty"`
	CreatedAt time.Time      `json:"created_at, omitempty"`
	UpdatedAt time.Time      `json:"updated_at, omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at, omitempty"`
}

func (a *Account) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	a.Secret = string(hashedPassword)
}
func (a *Account) Deposit(amount int) error {
	if amount > 0 {
		a.Balance += amount
		return nil
	}
	return fmt.Errorf("Deposit amount less than zero.")
}
func (a *Account) Transfer(amount int, accountDestination *Account) error {
	if amount <= a.Balance {
		a.Balance -= amount
		err := accountDestination.Deposit(amount)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("The originating account does not have sufficient balance.")
}
