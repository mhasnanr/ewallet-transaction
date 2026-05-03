package services

import (
	"context"

	"github.com/mhasnanr/ewallet-transaction/constants"
	"github.com/mhasnanr/ewallet-transaction/internal/models"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, int, *models.TransactionRequest) (*models.TransactionResponse, error)
	UpdateTransactionStatus(context.Context, string, string) error
}

type WalletAPI interface {
	CreditTransaction(ctx context.Context, userID int, request models.WalletRequest) error
	DebitTransaction(ctx context.Context, userID int, request models.WalletRequest) error
}

type TransactionService struct {
	repository TransactionRepository
	walletAPI  WalletAPI
}

func NewTransactionService(repository TransactionRepository, walletAPI WalletAPI) *TransactionService {
	return &TransactionService{repository, walletAPI}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, userID int, request *models.TransactionRequest) (*models.TransactionResponse, error) {
	var err error

	request.Status = constants.PendingTransaction
	newTransaction, err := s.repository.CreateTransaction(ctx, userID, request)
	if err != nil {
		return nil, err
	}

	var reference = newTransaction.Reference

	var externalWalletRequest = models.WalletRequest{
		Reference: reference,
		Amount:    request.Amount,
	}

	if request.Type == constants.PurchaseTransaction {
		err = s.walletAPI.DebitTransaction(ctx, userID, externalWalletRequest)
	} else {
		err = s.walletAPI.CreditTransaction(ctx, userID, externalWalletRequest)
	}

	if err != nil {
		errUpdateTransaction := s.repository.UpdateTransactionStatus(ctx, reference, constants.FailedTransaction)
		if errUpdateTransaction != nil {
			return nil, errUpdateTransaction
		}

		return nil, err
	}

	newTransaction.Status = constants.SuccessTransaction
	return newTransaction, s.repository.UpdateTransactionStatus(ctx, reference, constants.SuccessTransaction)
}
