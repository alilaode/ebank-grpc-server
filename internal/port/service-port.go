package port

import (
	"time"

	dbank "github.com/alilaode/ebank-grpc-server/internal/application/domain/bank"
	"github.com/google/uuid"
)

type BankServicePort interface {
	FindCurrentBalance(acct string) (float64, error)
	CreateExchangeRate(r dbank.ExchangeRate) (uuid.UUID, error)
	FindExchangeRate(fromCur string, toCur string, ts time.Time) (float64, error)
}
