package handler

import (
	"context"

	pb "github.com/mhasnanr/e-wallet/cmd/transaction"
	"github.com/mhasnanr/e-wallet/constants"
	"github.com/mhasnanr/e-wallet/internal/model"
)

type TransactionService interface {
	CreateTransaction(context.Context, *model.TransactionRequest) error
}

type Transaction struct {
	pb.UnimplementedTransactionServer
	svc TransactionService
}

func (t *Transaction) CreateTransaction(ctx context.Context, request *pb.TransactionRequest) error {
	transactionReq := &model.TransactionRequest{
		Amount:      float64(request.GetAmount()),
		Type:        request.GetType(),
		Description: request.Description,
		Status:      constants.PendingTransaction,
	}

	return t.svc.CreateTransaction(ctx, transactionReq)
}
