package constants

type TransactionType = string
type TransactionStatus = string

var (
	TopupTransaction    TransactionType = "TOP_UP"
	PurchaseTransaction TransactionType = "PURCHASE"
	RefundTransaction   TransactionType = "REFUND"
)

var (
	PendingTransaction TransactionStatus = "pending"
	FailedTransaction  TransactionStatus = "failed"
	SuccessTransaction TransactionStatus = "success"
)
