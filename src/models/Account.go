package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string
	Cpf     string `json:"password,omitempty"`
	Secret  string `json:"password,omitempty"`
	Balance int
}

func (a *Account) GetBalance() int {
	return int(a.Balance)
}
func (a *Account) GetPassword() string {
	return a.Secret
}
func (a *Account) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	a.Secret = string(hashedPassword)
}
func (a *Account) Deposit(amount int) (string, int) {
	if amount > 0 {
		a.Balance += amount
		return "Deposito realizado com sucesso!", int(a.Balance)
	}
	return "Valor do deposito menor que zero", int(a.Balance)
}
func (a *Account) Transfer(amount int, accountDestination *Account) (bool, error) {
	if amount > 0 && amount < a.Balance {
		a.Balance -= amount
		accountDestination.Deposit(amount)
		return true, nil
	}
	return false, fmt.Errorf("A conta de origem não tem saldo suficiente.")
}

var Accounts []Account
