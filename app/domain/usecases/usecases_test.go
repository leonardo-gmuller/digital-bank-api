package usecases_test

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/LeonardoMuller13/digital-bank-api/app/domain/usecases"
	"github.com/LeonardoMuller13/digital-bank-api/app/tests/mocks"
)

type MockSuit struct {
	ctrl             *gomock.Controller
	mockUow          *mocks.MockUowInterface
	mockAccountRepo  *mocks.MockAccountRepositoryInterface
	mockTransferRepo *mocks.MockTransfersRepositoryInterface
	useCase          *usecases.UseCase
}

func NewMockSuit(t *testing.T) *MockSuit {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockUow := mocks.NewMockUowInterface(ctrl)
	mockAccountRepo := mocks.NewMockAccountRepositoryInterface(ctrl)
	mockTransferRepo := mocks.NewMockTransfersRepositoryInterface(ctrl)

	useCase := &usecases.UseCase{
		Uow: mockUow,
	}

	return &MockSuit{
		ctrl:             ctrl,
		mockUow:          mockUow,
		mockAccountRepo:  mockAccountRepo,
		mockTransferRepo: mockTransferRepo,
		useCase:          useCase,
	}
}
