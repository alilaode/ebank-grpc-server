package application

import (
	"log"

	"github.com/alilaode/ebank-grpc-server/internal/port"
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
