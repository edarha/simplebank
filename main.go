package main

import (
	"database/sql"
	"log"

	"github.com/edarha/simplebank/api"
	db "github.com/edarha/simplebank/db/sqlc"
	"github.com/edarha/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	conf, err := util.LoadConfig(".")

	if err != nil {
		log.Panicf("Cannot load config %s", err)
	}

	util.Conf = conf

	conn, err := sql.Open(conf.Postgres.Driver, conf.Postgres.Source)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.Server.Address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
