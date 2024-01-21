package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const DBNAME = "sslchecker.db"

type SqliteStore struct {
	db *sql.DB
}

func NewSqliteUserStore() *SqliteStore {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Error in initializing db: %s", err)
	}
	return &SqliteStore{
		db: db,
	}
}

func initDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	if db, err = sql.Open("sqlite3", DBNAME); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			email TEXT,
			hashed_password TEXT,
			account_type TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS domains (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			domain TEXT,
			server_type TEXT,
			issuer TEXT,
			expires_in INTEGER,
			warn_before INTEGER,
			status INTEGER,
			last_checked TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
