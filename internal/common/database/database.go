package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	// Database is an interface to any db
	Database interface {
		Query(string, ...interface{}) (*sql.Rows, error)
		QueryRow(string, ...interface{}) *sql.Row
		Queryx(string, ...interface{}) (*sqlx.Rows, error)
		QueryRowx(string, ...interface{}) *sqlx.Row
		Exec(string, ...interface{}) (sql.Result, error)
		MustExec(string, ...interface{}) sql.Result
		Select(interface{}, string, ...interface{}) error
		Get(interface{}, string, ...interface{}) error
		NamedExec(string, interface{}) (sql.Result, error)
		Begin() (*sql.Tx, error)
		Beginx() (*sqlx.Tx, error)
		MustBegin() *sqlx.Tx
		Rebind(string) string
		Ping() error
	}
)

// Init will init database connection based on selected db
func Init(dbVal, dbURI string) (dbObj Database) {
	// we can add more db implementation
	switch dbVal {
	case "sqlx":
		dbObj = InitSqlx(dbURI)
	default:
		dbObj = InitSqlx(dbURI)
	}

	return
}
