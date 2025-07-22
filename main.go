package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yogesh-k64/middleware-simple-bank/api"
	db "github.com/yogesh-k64/middleware-simple-bank/db/sqlc"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

func main() {

	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config %v\n", err)
	}

	dbCon, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(dbCon)

	server, err := api.NewServer(store, config)

	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
