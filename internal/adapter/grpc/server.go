package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/alilaode/ebank-grpc-proto/protogen/go/bank"
	"github.com/alilaode/ebank-grpc-server/internal/port"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	bankService port.BankServicePort
	grpcPort    int
	server      *grpc.Server
	bank.BankServiceServer
}

func NewGrpcAdapter(
	bankService port.BankServicePort,
	grpcPort int,
) *GrpcAdapter {
	return &GrpcAdapter{
		bankService: bankService,
		grpcPort:    grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	// creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")

	// if err != nil {
	// 	log.Fatalln("Can't create server credentials :", err)
	// }

	grpcServer := grpc.NewServer(
	// grpc.Creds(creds),
	// grpc.ChainUnaryInterceptor(
	// 	interceptor.LogUnaryServerInterceptor(),
	// 	interceptor.BasicUnaryServerInterceptor(),
	// ),
	// grpc.ChainStreamInterceptor(
	// 	interceptor.LogStreamServerInterceptor(),
	// 	interceptor.BasicStreamServerInterceptor(),
	// ),
	)

	a.server = grpcServer
	bank.RegisterBankServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
