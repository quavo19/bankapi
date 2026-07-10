package main

import (
	"log"
	"database/sql"
	db "github.com/quavo19/bankapi/db/sqlc"
	"github.com/quavo19/bankapi/api"
	_ "github.com/lib/pq"
)

const (
	dbDriver        = "postgres"
	defaultDBSource = "postgresql://root:07May2002@@localhost:5433/simple_bank?sslmode=disable"
	serverAddress   = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, defaultDBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}