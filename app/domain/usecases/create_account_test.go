package usecases_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/entity"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/usecases"
	"github.com/leonardo-gmuller/digital-bank-api/app/pkg/uow"
)

func TestCreateAccount_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	input := usecases.InputCreateAccount{
		Name:    "John Doe",
		Cpf:     "20652942059",
		Secret:  "senhaSegura",
		Balance: 1000,
	}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.Cpf).Return(entity.Account{}, erring.ErrAccountExists).Times(1)
	suit.mockAccountRepo.EXPECT().Create(gomock.Any(), input.Name, input.Cpf, gomock.Any(), gomock.Any()).Return(nil).Times(1)

	err := suit.useCase.CreateAccount(context.Background(), input)
	assert.NoError(t, err, "Unexpected error on created account")
}

func TestCreateAccount_ValidationFailure(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	input := usecases.InputCreateAccount{
		Name:    "John Doe",
		Cpf:     "123",
		Secret:  "password123",
		Balance: 1000,
	}

	err := suit.useCase.CreateAccount(context.Background(), input)
	assert.ErrorIs(t, err, erring.ErrAccountCPFIsInvalid)
}

func TestCreateAccount_AccountAlreadyExists(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	input := usecases.InputCreateAccount{
		Name:    "John Doe",
		Cpf:     "20652942059",
		Secret:  "password123",
		Balance: 1000,
	}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.Cpf).Return(entity.Account{}, nil).Times(1)
	err := suit.useCase.CreateAccount(context.Background(), input)

	assert.ErrorIs(t, err, erring.ErrAccountExists)
}

func TestCreateAccount_DatabaseFailure(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	input := usecases.InputCreateAccount{
		Name:    "John Doe",
		Cpf:     "20652942059",
		Secret:  "password123",
		Balance: 1000,
	}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.Cpf).Return(entity.Account{}, erring.ErrAccountExists).Times(1)
	suit.mockAccountRepo.EXPECT().Create(gomock.Any(), input.Name, input.Cpf, gomock.Any(), gomock.Any()).Return(errors.New("database error")).Times(1)

	err := suit.useCase.CreateAccount(context.Background(), input)

	assert.Error(t, err, err.Error())
}
