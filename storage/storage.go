package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"rtArchive/config"
)

type (
	DBS struct {
		pgSQL *sql.DB
	}
)

func (dbs *DBS) Connect() (err error) {
	if dbs.pgSQL, err = sql.Open("postgres", config.PgSqlDSN); err != nil {
		return err
	}
	if err = dbs.pgSQL.Ping(); err != nil {
		return err
	}
	return nil
}
