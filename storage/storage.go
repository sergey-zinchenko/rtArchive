package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type (
	DBS struct {
		pgSQL *sql.DB
	}
)

func (dbs *DBS) Connect() (err error) {
	//connStr := fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", config.DBUserName, config.DBName)
	//fmt.Println(connStr)
	if dbs.pgSQL, err = sql.Open("postgres", ""); err != nil {
		log.Fatal(err)
	}
	return nil
}
