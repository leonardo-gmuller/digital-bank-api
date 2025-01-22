package usecases_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/dto"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/entity"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/erring"
	"github.com/LeonardoMuller13/digital-bank-api/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/app/pkg/uow"
)

var input = usecases.InputCreateTransfer{
	AccountDestinationCPF: "12345678900",
	Amount:                500,
}

func TestCreateTransfer_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	accountOrigin := entity.Account{ID: 1, Cpf: "11111111111", Balance: 1000}
	accountDest := entity.Account{ID: 2, Cpf: input.AccountDestinationCPF, Balance: 2000}

	ctx := context.WithValue(context.Background(), dto.UserKey, &dto.User{ID: int(accountOrigin.ID)})

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil)
	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)

	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.AccountDestinationCPF).Return(accountDest, nil).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(int64(accountOrigin.ID))).Return(accountOrigin, nil).Times(1)

	suit.mockTransferRepo.EXPECT().Create(gomock.Any(), gomock.Eq(int(accountOrigin.ID)), gomock.Eq(int(accountDest.ID)), input.Amount).Return(nil).Times(1)

	suit.mockAccountRepo.EXPECT().UpdateBalance(gomock.Any(), gomock.Eq(int(accountOrigin.ID)), gomock.Any()).Return(nil).Times(1)
	suit.mockAccountRepo.EXPECT().UpdateBalance(gomock.Any(), gomock.Eq(int(accountDest.ID)), gomock.Any()).Return(nil).Times(1)

	err := suit.useCase.CreateTransfer(ctx, input)

	assert.NoError(t, err, "Unexpected error on created transfer")
}

func TestCreateTransfer_DestinationAccountNotFound(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	accountDest := entity.Account{Cpf: input.AccountDestinationCPF}

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil)
	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), accountDest.Cpf).Return(entity.Account{}, errors.New("not found")).Times(1)

	err := suit.useCase.CreateTransfer(context.Background(), input)

	require.Error(t, err)
	assert.ErrorIs(t, err, erring.ErrTransferAccountDestinationNotFound)
}

func TestCreateTransfer_UserNotFound(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	accountOrigin := entity.Account{ID: 1}
	accountDest := entity.Account{ID: 2, Cpf: input.AccountDestinationCPF, Balance: 2000}

	ctx := context.WithValue(context.Background(), dto.UserKey, &dto.User{ID: int(accountOrigin.ID)})

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil)
	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.AccountDestinationCPF).Return(accountDest, nil).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(int64(accountOrigin.ID))).Return(entity.Account{}, errors.New("user not found")).Times(1)

	err := suit.useCase.CreateTransfer(ctx, input)

	require.Error(t, err)
	assert.ErrorIs(t, err, erring.ErrTransferUserNotFound)
}

func TestCreateTransfer_FailedBalanceUpdate(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	accountOrigin := entity.Account{ID: 1, Cpf: "11111111111", Balance: 1000}
	accountDest := entity.Account{ID: 2, Cpf: input.AccountDestinationCPF, Balance: 2000}

	ctx := context.WithValue(context.Background(), dto.UserKey, &dto.User{ID: int(accountOrigin.ID)})

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil)
	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "AccountRepository").Return(suit.mockAccountRepo, nil)
	suit.mockUow.EXPECT().Do(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, fn func(uow *uow.Uow) error) error {
		return fn(&uow.Uow{})
	}).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByCpf(gomock.Any(), input.AccountDestinationCPF).Return(accountDest, nil).Times(1)
	suit.mockAccountRepo.EXPECT().GetAccountByID(gomock.Any(), gomock.Eq(int64(accountOrigin.ID))).Return(accountOrigin, nil).Times(1)

	suit.mockTransferRepo.EXPECT().Create(gomock.Any(), gomock.Eq(int(accountOrigin.ID)), gomock.Eq(int(accountDest.ID)), input.Amount).Return(nil).Times(1)

	suit.mockAccountRepo.EXPECT().UpdateBalance(gomock.Any(), gomock.Eq(int(accountOrigin.ID)), gomock.Any()).Return(errors.New("failed update balance")).Times(1)

	err := suit.useCase.CreateTransfer(ctx, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update balance for account id 1")
}
