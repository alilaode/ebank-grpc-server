package grpc

import (
	"context"
	"log"
	"time"

	"github.com/alilaode/ebank-grpc-proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context,
	req *bank.CurrentBalanceRequest) (*bank.CurrentBalanceResponse, error) {
	now := time.Now()
	bal, err := a.bankService.FindCurrentBalance(req.AccountNumber)

	if err != nil {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"account %v not found", req.AccountNumber,
		)
	}

	return &bank.CurrentBalanceResponse{
		Amount: bal,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}

func (a *GrpcAdapter) FetchExchangeRates(req *bank.ExchangeRateRequest,
	stream bank.BankService_FetchExchangeRatesServer) error {
	context := stream.Context()

	for {
		select {
		case <-context.Done():
			log.Println("Client cancelled stream")
			return nil
		default:
			now := time.Now().Truncate(time.Second)
			rate, _ := a.bankService.FindExchangeRate(req.FromCurrency, req.ToCurrency, now)

			stream.Send(
				&bank.ExchangeRateResponse{
					FromCurrency: req.FromCurrency,
					ToCurrency:   req.ToCurrency,
					Rate:         rate,
					Timestamp:    now.Format(time.RFC3339),
				},
			)

			log.Printf("Exchange rate sent to client, %v to %v : %v\n", req.FromCurrency,
				req.ToCurrency, rate)

			time.Sleep(3 * time.Second)

		}
	}

}
