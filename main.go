package main

import (
	"database/sql"
	"log"

	"github.com/hiamthach/golang-udemy-backend/api"
	db "github.com/hiamthach/golang-udemy-backend/db/sqlc"
	"github.com/hiamthach/golang-udemy-backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	var conn *sql.DB
	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
