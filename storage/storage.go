package storage

import (
	"database/sql"
	"fmt"
	"github.com/kataras/iris/core/errors"
	// empty inport pq package
	_ "github.com/lib/pq"
	"rtArchive/config"
	"rtArchive/storage/queries"
)

//DBS struct provide connection to db
type (
	DBS struct {
		pgSQL *sql.DB
	}
)

//Connect - connect to db
func (dbs *DBS) Connect() (err error) {
	if dbs.pgSQL, err = sql.Open("postgres", config.PgSqlDSN); err != nil {
		return err
	}
	if err = dbs.pgSQL.Ping(); err != nil {
		return err
	}
	return nil
}

//CheckTables - transaction for db instantiation/check
func (dbs *DBS) CheckTables() (err error) {
	if dbs.pgSQL == nil {
		return errors.New("pgSql database not connected")
	}
	var transaction = func(tx *sql.Tx) error {
		if _, err = tx.Exec(queries.CreateSchema); err != nil {
			fmt.Println("1", err)
			return err
		}
		if _, err = tx.Exec(queries.CreateDialogSourceEnum); err != nil {
			fmt.Println("2", err)
			return err
		}
		if _, err = tx.Exec(queries.CreateTableRoundTrips); err != nil {
			fmt.Println("3", err)
			return err
		}
		if _, err = tx.Exec(queries.CreateUpdatedAtFunction); err != nil {
			fmt.Println("4", err)
			return err
		}
		if _, err = tx.Exec(queries.CreateUpdatedAtTrigger); err != nil {
			fmt.Println("5", err)
			return err
		}
		return nil
	}
	tx, err := dbs.pgSQL.Begin()
	if err != nil {
		return err
	}
	err = transaction(tx)
	if err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
