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
}
