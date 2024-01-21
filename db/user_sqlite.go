package db

import "github.com/AshirwadPradhan/tracksslcerts/types"

func (us *SqliteStore) CreateUser(user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	INSERT INTO users 
	(username, email, hashed_password, account_type)
	VALUES (?, ?, ?, ?)
	`, user.UserName, user.Email, user.HashedPassword, user.AccountType)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) UpdateUserUsername(username string, user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE users SET 
	username = ? WHERE username = ?
	`, username, user.UserName)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) UpdateUserPassword(password string, user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE users SET 
	hashed_password = ? WHERE username = ?
	`, password, user.UserName)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) UpdateUserAccountType(accountType string, user *types.User) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE users SET 
	account_type = ? WHERE username = ?
	`, accountType, user.UserName)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func (us *SqliteStore) DeleteUser(user *types.User) error {
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

func (us *SqliteStore) ValidatePassword(password string, user *types.User) bool {
	var hashedPassword string

	tx, err := us.db.Begin()
	if err != nil {
		return false
	}

	rows, err := tx.Query(`
	SELECT hashed_password FROM users WHERE username = ?
	`, user.UserName)

	if err != nil {
		tx.Rollback()
		return false
	}
	defer rows.Close()

	err = rows.Scan(hashedPassword)
	if err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return hashedPassword == password
}

func (us *SqliteStore) ReadUser(username string) (*types.User, error) {
	user := types.User{}

	tx, err := us.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(`
	SELECT username, email, account_type 
	FROM users WHERE username = ?
	`, username)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.UserName, &user.Email, &user.AccountType)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &user, nil
}
