package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mhasnanr/ewallet-transaction/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, userID int, request *models.TransactionRequest) (*models.TransactionResponse, error) {
	reference := fmt.Sprintf("TRX-%d", time.Now().UnixNano())

	transaction := models.Transaction{
		UserID:      userID,
		Amount:      request.Amount,
		Type:        request.Type,
		Status:      request.Status,
		Reference:   reference,
		Description: request.Description,
	}

	if err := r.db.WithContext(ctx).Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &models.TransactionResponse{
		ID:        transaction.ID,
		Type:      transaction.Type,
		Reference: transaction.Reference,
		Status:    transaction.Status,
	}, nil
}

func (r *TransactionRepository) UpdateTransactionStatus(ctx context.Context, reference string, status string) error {
	return r.db.WithContext(ctx).Model(&models.Transaction{}).
		Where("reference = ?", reference).
		Update("status", status).Error
}
