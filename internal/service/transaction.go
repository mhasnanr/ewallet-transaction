package service

import (
	"context"

	"github.com/mhasnanr/e-wallet/constants"
	"github.com/mhasnanr/e-wallet/internal/model"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, int, *model.TransactionRequest) (model.TransactionResponse, error)
	UpdateTransactionStatus(context.Context, string, string) error
}

type WalletGRPC interface {
	CreditTransaction(ctx context.Context, userID int, request model.ExternalTransactionRequest) error
	DebitTransaction(ctx context.Context, userID int, request model.ExternalTransactionRequest) error
}

type TransactionService struct {
	repository TransactionRepository
	walletSvc  WalletGRPC
}

func NewTransactionService(repository TransactionRepository, walletGRPC WalletGRPC) *TransactionService {
	return &TransactionService{repository, walletGRPC}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, userID int, request *model.TransactionRequest) error {
	var err error
	newTransaction, err := s.repository.CreateTransaction(ctx, userID, request)
	if err != nil {
		return err
	}

	var reference = newTransaction.Reference

	var externalWalletRequest = model.ExternalTransactionRequest{
		Reference: reference,
		Amount:    request.Amount,
	}

	if request.Type == constants.PurchaseTransaction {
		err = s.walletSvc.CreditTransaction(ctx, userID, externalWalletRequest)
	} else {
		err = s.walletSvc.DebitTransaction(ctx, userID, externalWalletRequest)
	}

	if err != nil {
		errUpdateTransaction := s.repository.UpdateTransactionStatus(ctx, reference, constants.FailedTransaction)
		if errUpdateTransaction != nil {
			return errUpdateTransaction
		}

		return err
	}

	return s.repository.UpdateTransactionStatus(ctx, reference, constants.SuccessTransaction)
}
