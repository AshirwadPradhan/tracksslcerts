package db

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/AshirwadPradhan/tracksslcerts/types"
)

const DBNAME = "sslchecker.db"

type SqliteUserStore struct {
	db *sql.DB
}

func NewSqliteUserStore() *SqliteUserStore {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Error in initializing db: %s", err)
	}
	return &SqliteUserStore{
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
			account_type TEXT,
			tracked_domains TEXT
		)
	`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (us *SqliteUserStore) Create(user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO users 
	(username, email, hashed_password, account_type, tracked_domains)
	VALUES (?, ?, ?, ?, ?)`,
		user.UserName, user.Email, user.HashedPassword, user.AccountType, "")

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *SqliteUserStore) Update(user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	// Serialize tracked domains to JSON
	trackedDomains, err := json.Marshal(user.TrackedDomains)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec(`
	UPDATE users SET 
	username = ?, email = ?, hashed_password = ?, account_type = ?, tracked_domains = ? 
	WHERE username = ?
	`, user.UserName, user.Email, user.HashedPassword, user.AccountType, string(trackedDomains), user.UserName)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (us *SqliteUserStore) Delete(user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	DELETE from users WHERE username = ?
	`, user.UserName)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *SqliteUserStore) Read(username string) (*types.User, error) {
	user := types.User{}

	tx, err := us.db.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select username, email, account_type, tracked_domains from users where username = ?", username)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	var trackedDomains string
	err = rows.Scan(&user.UserName, &user.Email, &user.AccountType, &trackedDomains)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = json.Unmarshal([]byte(trackedDomains), &user.TrackedDomains)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &user, nil
}
