package data

import (
	"database/sql"
	"os"
	"utils/apperrors"
)

var db *sql.DB
var stmt *sql.Stmt

func init() {
	var err error
	dbHandler := os.Getenv("DB_HANDLER")
	dbDriver := os.Getenv("DB_DRIVER")
	db, err = sql.Open(dbDriver, dbHandler)
	apperrors.ExitOnError(err)
}

/*
GetDbHandler returns
the db handler object
run queries against this
*/
func GetDbHandler() *sql.DB {
	return db
}
