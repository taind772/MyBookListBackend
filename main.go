package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ndtai772/MyBookListBackend/api"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	// sampledata "github.com/ndtai772/MyBookListBackend/sample_data"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	// sampledata.Gen(store)

	server := api.NewServer(store)

	server.Start(serverAddress)
}
