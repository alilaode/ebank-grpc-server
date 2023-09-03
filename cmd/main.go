package main

import (
	mygrpc "github.com/alilaode/ebank-grpc-server/internal/adapter/grpc"
	app "github.com/alilaode/ebank-grpc-server/internal/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	bs := &app.BankService{}

	grpcAdapter := mygrpc.NewGrpcAdapter(bs, 9090)

	grpcAdapter.Run()

}
