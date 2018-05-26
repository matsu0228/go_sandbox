package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DB is variable
type DB *sqlx.DB

// NewDB is initialize database connection
func NewDB(server, port, user, password, database string) (DB, error) {
	db, err := connectDatabase(server, port, user, password, database)
	// if err != nil {
	// 	nildb := sqlx.DB
	// 	return *nildb, err
	// }
	return db, err
}

// Close is closer of database
//  -> Why cant use followings..?
// func (db DB) Close() {
// 	db.Close()
// }

func connectDatabase(server, port, user, password, database string) (*sqlx.DB, error) {
	// Build connection string
	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database,
	)

	db, err := sqlx.Connect("sqlserver", connString)
	if err != nil {
		return new(sqlx.DB), err
	}
	return db, err
}
