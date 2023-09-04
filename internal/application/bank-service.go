package application

import (
	"log"
	"time"

	db "github.com/alilaode/ebank-grpc-server/internal/adapter/database"
	dbank "github.com/alilaode/ebank-grpc-server/internal/application/domain/bank"
	"github.com/alilaode/ebank-grpc-server/internal/port"
	"github.com/google/uuid"
)

type BankService struct {
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}

func (s *BankService) FindCurrentBalance(acct string) (float64, error) {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		log.Println("Error on FindCurrentBalance :", err)
		return 0, err
	}

	return bankAccount.CurrentBalance, nil
}

func (s *BankService) CreateExchangeRate(r dbank.ExchangeRate) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	exchangeRateOrm := db.BankExchangeRateOrm{
		ExchangeRateUuid:   newUuid,
		FromCurrency:       r.FromCurrency,
		ToCurrency:         r.ToCurrency,
		Rate:               r.Rate,
		ValidFromTimestamp: r.ValidFromTimestamp,
		ValidToTimestamp:   r.ValidToTimestamp,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	return s.db.CreateExchangeRate(exchangeRateOrm)
}

func (s *BankService) FindExchangeRate(fromCur string, toCur string, ts time.Time) (float64, error) {
	exchangeRate, err := s.db.GetExchangeRateAtTimestamp(fromCur, toCur, ts)

	if err != nil {
		return 0, err
	}

	return float64(exchangeRate.Rate), nil
}
