package models

type WalletRequest struct {
	UserID    int
	Amount    int64
	Reference string
}

type WalletResponse struct {
	Balance int64
}
