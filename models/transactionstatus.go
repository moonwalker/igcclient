package models

type TransactionStatus int

const (
	TS_Requested            TransactionStatus = 1
	TS_Approved             TransactionStatus = 2
	TS_Rejected             TransactionStatus = 3
	TS_Cancelled            TransactionStatus = 4
	TS_Error                TransactionStatus = 5
	TS_WaitingApproval      TransactionStatus = 6
	TS_Pending              TransactionStatus = 7
	TS_BankProcessing       TransactionStatus = 8
	TS_Reversal             TransactionStatus = 10
	TS_CancelWithdraw       TransactionStatus = 11
	TS_WithdrawRejected     TransactionStatus = 12
	TS_PendingManualDeposit TransactionStatus = 13
)
