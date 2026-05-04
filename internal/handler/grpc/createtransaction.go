package grpc

import (
	"context"

	pb "github.com/mhasnanr/ewallet-transaction/cmd/transaction"
	"github.com/mhasnanr/ewallet-transaction/constants"
	"github.com/mhasnanr/ewallet-transaction/internal/models"
)

type TransactionService interface {
	CreateTransaction(context.Context, int, *models.TransactionRequest) (*models.TransactionResponse, error)
}

type Transaction struct {
	pb.UnimplementedTransactionServer
	svc TransactionService
}

func NewTransactionHandler(svc TransactionService) *Transaction {
	return &Transaction{
		svc: svc,
	}
}

func (t *Transaction) CreateTransaction(ctx context.Context, request *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	transactionReq := &models.TransactionRequest{
		Amount:      request.GetAmount(),
		Type:        request.GetType(),
		Description: request.GetDescription(),
	}

	res, err := t.svc.CreateTransaction(ctx, int(request.GetUserID()), transactionReq)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Message: constants.TransactionSuccess,
		Data: &pb.TransactionData{
			Status:    res.Status,
			Reference: res.Reference,
			Type:      res.Type,
		},
	}, nil
}
