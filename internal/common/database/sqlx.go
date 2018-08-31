package database

import (
	"github.com/nabillarahmani/currencyapp/internal/common/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	//SqlxDB wrapper based on sqlt
	SqlxDB struct {
		db *sqlx.DB
	}
)

// InitSqlx is an init function
func InitSqlx(dbURI string) (db Database) {
	db, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Error(err, "Error when connect")
		log.Fatal("error when connect to db")
	}

	return
}
