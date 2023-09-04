package port

import (
	db "github.com/alilaode/ebank-grpc-server/internal/adapter/database"
)

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(acct string) (db.BankAccountOrm, error)
}
