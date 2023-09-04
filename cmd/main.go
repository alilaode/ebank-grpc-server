package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	dbmigration "github.com/alilaode/ebank-grpc-server/db"
	_ "github.com/jackc/pgx/v5/stdlib"

	mydb "github.com/alilaode/ebank-grpc-server/internal/adapter/database"
	mygrpc "github.com/alilaode/ebank-grpc-server/internal/adapter/grpc"
	app "github.com/alilaode/ebank-grpc-server/internal/application"
	"github.com/alilaode/ebank-grpc-server/internal/application/domain/bank"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://root:secret@localhost:5432/grpc?sslmode=disable")

	if err != nil {
		log.Fatalln("Can't connect database :", err)
	}

	dbmigration.Migrate(sqlDB)

	databaseAdapter, err := mydb.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter :", err)
	}

	bs := app.NewBankService(databaseAdapter)

	// generare ExchangeRates
	go generateExchangeRates(bs, "USD", "IDR", 5*time.Second)

	grpcAdapter := mygrpc.NewGrpcAdapter(bs, 9090)

	grpcAdapter.Run()

}

func generateExchangeRates(bs *app.BankService, fromCurrency, toCurrency string, duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(duration).Add(-1 * time.Millisecond)

		dummyRate := bank.ExchangeRate{
			FromCurrency:       fromCurrency,
			ToCurrency:         toCurrency,
			ValidFromTimestamp: validFrom,
			ValidToTimestamp:   validTo,
			Rate:               2000 + float64(rand.Intn(300)),
		}

		bs.CreateExchangeRate(dummyRate)
	}
}
