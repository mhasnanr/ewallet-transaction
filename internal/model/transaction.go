package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID        int     `gorm:"primaryKey"`
	UserID    int     `gorm:"user_id"`
	Amount    float64 `gorm:"column:amount"`
	Type      string  `gorm:"column:type"`
	Status    string  `gorm:"column:status"`
	Reference string  `gorm:"column:reference"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*Transaction) TableName() string {
	return "transactions"
}

type TransactionRequest struct {
	Amount      float64 `json:"amount" gorm:"column:amount;type:decimal(15, 2)" validate:"required,gt=0"`
	Type        string  `json:"type" gorm:"column:type" validate:"required"`
	Description string  `json:"description" gorm:"column:description"`
	Status      string  `json:"status" gorm:"column:status" validate:"required"`
}

func (f *TransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(f)
}

type TransactionResponse struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Reference string `json:"description"`
	Status    string `json:"status"`
}

type UpdateTransaction struct {
	Status *string `gorm:"column:status"`
}

type ExternalTransactionRequest struct {
	Reference string
	Amount    float64
}
