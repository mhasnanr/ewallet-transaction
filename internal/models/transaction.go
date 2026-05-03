package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID          int    `gorm:"primaryKey"`
	UserID      int    `gorm:"user_id"`
	Amount      int64  `gorm:"column:amount"`
	Type        string `gorm:"column:type"`
	Status      string `gorm:"column:status"`
	Reference   string `gorm:"column:reference"`
	Description string `gorm:"column:description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (*Transaction) TableName() string {
	return "transactions"
}

type TransactionRequest struct {
	Amount      int64  `json:"amount" gorm:"column:amount;type:bigint" validate:"required,gt=0"`
	Type        string `json:"type" gorm:"column:type" validate:"required"`
	Description string `json:"description" gorm:"column:description"`
	Status      string
}

func (f *TransactionRequest) Validate() error {
	v := validator.New()
	return v.Struct(f)
}

type TransactionResponse struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

type UpdateTransaction struct {
	Status *string `gorm:"column:status"`
}
