package usecases_test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/leonardo-gmuller/digital-bank-api/app/domain/dto"
)

func TestListUserTransfer_Success(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	expectedTransfers := []dto.OutputTransfer{
		{AccountDestinationCPF: "12345678900", Amount: 100},
		{AccountDestinationCPF: "09876543211", Amount: 200},
	}

	ctx := context.WithValue(context.Background(), dto.UserKey, &dto.User{ID: int(1)})

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil)
	suit.mockTransferRepo.EXPECT().List(gomock.Any(), 1).Return(expectedTransfers, nil)
	suit.mockUow.EXPECT().CommitOrRollback(gomock.Any()).Return(nil)

	output, err := suit.useCase.ListUserTransfer(ctx)
	require.NoError(t, err)
	assert.Len(t, output.Transfers, 2)
	assert.Equal(t, expectedTransfers, output.Transfers)
}

func TestListUserTransfer_FailUserContext(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	ctx := context.Background()

	_, err := suit.useCase.ListUserTransfer(ctx)

	require.Error(t, err)
	assert.Equal(t, "failed to assert context value as dto.User", err.Error())
}

func TestListUserTransfer_FailRepositoryError(t *testing.T) {
	t.Parallel()
	suit := NewMockSuit(t)

	defer suit.ctrl.Finish()

	mockUser := &dto.User{ID: 1}
	ctx := context.WithValue(context.Background(), dto.UserKey, mockUser)

	suit.mockUow.EXPECT().GetRepository(gomock.Any(), "TransferRepository").Return(suit.mockTransferRepo, nil).Times(1)
	suit.mockTransferRepo.EXPECT().List(ctx, mockUser.ID).Return(nil, errors.New("database error")).Times(1)
	suit.mockUow.EXPECT().CommitOrRollback(ctx).Return(nil).Times(1)

	_, err := suit.useCase.ListUserTransfer(ctx)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list transfers for id")
}
