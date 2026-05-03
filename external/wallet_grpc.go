package external

import (
	"context"
	"errors"

	"github.com/mhasnanr/ewallet-transaction/bootstrap"
	pb "github.com/mhasnanr/ewallet-transaction/cmd/wallet"
	"github.com/mhasnanr/ewallet-transaction/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WalletGRPC struct {
	client pb.WalletClient
}

func NewWalletGRPC() (*WalletGRPC, *grpc.ClientConn, error) {
	serverAddr := bootstrap.GetEnv("WALLET_GRPC_URL", "")

	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, errors.New("failed to dial ums grpc")
	}

	client := pb.NewWalletClient(conn)

	return &WalletGRPC{
		client: client,
	}, conn, nil
}

func (e *WalletGRPC) DebitTransaction(ctx context.Context, userID int, req models.WalletRequest) error {
	request := &pb.WalletRequest{
		UserID:    int64(userID),
		Amount:    int64(req.Amount),
		Reference: req.Reference,
	}

	_, err := e.client.DebitBalance(ctx, request)
	if err != nil {
		return errors.New("failed to perform debit balance")
	}

	return nil
}

func (e *WalletGRPC) CreditTransaction(ctx context.Context, userID int, req models.WalletRequest) error {
	request := &pb.WalletRequest{
		UserID:    int64(userID),
		Amount:    int64(req.Amount),
		Reference: req.Reference,
	}

	_, err := e.client.CreditBalance(ctx, request)
	if err != nil {
		return errors.New("failed to perform credit balance")
	}

	return nil
}
