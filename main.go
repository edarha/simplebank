package main

import (
	"database/sql"
	"log"

	"github.com/edarha/simplebank/api"
	db "github.com/edarha/simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	address  = ":8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
