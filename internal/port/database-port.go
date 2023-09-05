package port

import (
	"time"

	"github.com/google/uuid"

	db "github.com/alilaode/ebank-grpc-server/internal/adapter/database"
)

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(acct string) (db.BankAccountOrm, error)
	CreateExchangeRate(r db.BankExchangeRateOrm) (uuid.UUID, error)
	GetExchangeRateAtTimestamp(fromCur string, toCur string, ts time.Time) (db.BankExchangeRateOrm, error)
	CreateTransaction(acct db.BankAccountOrm, t db.BankTransactionOrm) (uuid.UUID, error)
	CreateTransfer(transfer db.BankTransferOrm) (uuid.UUID, error)
	CreateTransferTransactionPair(fromAccountOrm db.BankAccountOrm, toAccountOrm db.BankAccountOrm,
		fromTransactionOrm db.BankTransactionOrm, toTransactionOrm db.BankTransactionOrm) (bool, error)
	UpdateTransferStatus(transfer db.BankTransferOrm, status bool) error
}
