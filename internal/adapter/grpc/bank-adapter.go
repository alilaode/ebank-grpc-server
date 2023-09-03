package grpc

import (
	"context"
	"time"

	"github.com/alilaode/ebank-grpc-proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
)

func (a *GrpcAdapter) GetCurrentBalance(ctx context.Context, req *bank.CurrentBalanceRequest) (res *bank.CurrentBalanceResponse, err error) {
	now := time.Now()

	bal := a.bankService.FindCurrentBalance(req.AccountNumber)

	return &bank.CurrentBalanceResponse{
		Amount: bal,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}
