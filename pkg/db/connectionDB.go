package db

import (
	"database/sql"
)

func DB() (*sql.DB, error) {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root:21032991@tcp(localhost:3306)/products")
	if err != nil {
		return nil, err
	}
	return db, nil
}
