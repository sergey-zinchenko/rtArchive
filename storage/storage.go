package storage

import (
	"database/sql"
	"fmt"
	"log"
	"rtArchive/config"
)

type (
	DBS struct {
		pgSQL *sql.DB
	}
)

func (dbs *DBS) Connect() error {
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", config.DBUserName, config.DBName)
	dbs.pgSQL, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
