package main

import (
	"database/sql"
	"log"

	"github.com/judgegodwins/simplebank/api"
	db "github.com/judgegodwins/simplebank/db/sqlc"
	"github.com/judgegodwins/simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(dbConn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
