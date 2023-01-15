package main

import (
	"database/sql"
	"log"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	util "simple_bank/db/utils"

	_ "github.com/lib/pq"
)



func main() {
	config,err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("connect load config ", err)
	}

	conn, err := sql.Open(config.DBDRiver, config.DBSource)

	if err != nil {
		log.Fatal("connot connet to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil{
		log.Fatal("Error to Start Server", err)
	}
}