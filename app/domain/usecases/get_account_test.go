package usecases_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/entity"
	"github.com/leonardo-gmuller/digital-bank-api/app/domain/erring"
)

func TestListAccount_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	expectedAccounts := []dto.ResponseAccount{
		{ID: 1, Name: "John Doe", Cpf: "12345678900"},
		{ID: 2, Name: "Jane Doe", Cpf: "09876543211"},
	}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().List(gomock.Any()).Return(expectedAccounts, nil)
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	output, err := suit.useCase.ListAccount(context.Background())

	require.NoError(t, err)
	assert.Len(t, output.Accounts, 2)
	assert.Equal(t, expectedAccounts, output.Accounts)
}

func TestListAccount_Failure(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("database error"))
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	_, err := suit.useCase.ListAccount(context.Background())

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list accounts")
}

func TestGetBalanceByID_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	expectedBalance := dto.ResponseAccountBalance{Balance: 1500}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().GetAccountBalanceByID(gomock.Any(), int64(1)).Return(expectedBalance, nil)
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	result, err := suit.useCase.GetBalanceByID(context.Background(), 1)

	require.NoError(t, err)
	assert.Equal(t, expectedBalance.Balance, result.Balance)
}

func TestGetBalanceByID_Failure(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().GetAccountBalanceByID(gomock.Any(), int64(1)).Return(dto.ResponseAccountBalance{}, errors.New("not found"))
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	_, err := suit.useCase.GetBalanceByID(context.Background(), 1)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get account balance")
}

func TestGetAccountByCpf_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	expectedAccount := entity.Account{
		Name:   "John Doe",
		Cpf:    "12345678900",
		Secret: "hashedpassword",
	}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), "12345678900").Return(expectedAccount, nil)
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	result, err := suit.useCase.GetAccountByCpf(context.Background(), "12345678900")

	require.NoError(t, err)
	assert.Equal(t, expectedAccount.Name, result.Name)
	assert.Equal(t, expectedAccount.Cpf, result.Cpf)
}

func TestGetAccountByCpf_Failure(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), "12345678900").Return(entity.Account{}, errors.New("not found"))
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	_, err := suit.useCase.GetAccountByCpf(context.Background(), "12345678900")

	require.Error(t, err)
	assert.Equal(t, erring.ErrAccountNotExists, err)
}
