package main

import (
	"database/sql"
	"log"

	dbmigration "github.com/alilaode/ebank-grpc-server/db"
	_ "github.com/jackc/pgx/v5/stdlib"

	mydb "github.com/alilaode/ebank-grpc-server/internal/adapter/database"
	mygrpc "github.com/alilaode/ebank-grpc-server/internal/adapter/grpc"
	app "github.com/alilaode/ebank-grpc-server/internal/application"
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

	grpcAdapter := mygrpc.NewGrpcAdapter(bs, 9090)

	grpcAdapter.Run()

}
