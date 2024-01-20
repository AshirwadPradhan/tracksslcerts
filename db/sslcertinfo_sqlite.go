package db

import (
	"github.com/AshirwadPradhan/tracksslcerts/types"
)

func (us *SqliteStore) AddDomainsToTrack(sslCertInfo []types.SSLCertInfo) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	for _, cert := range sslCertInfo {
		_, err = tx.Exec(`
		INSERT INTO domains 
		(username, domain, server_type, issuer, expires_in, warn_before, status, last_checked)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`, cert.UserName, cert.Domain, cert.ServerType, cert.Issuer, cert.ExpiresIn, cert.WarnBefore, cert.Status, cert.LastChecked)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) DeleteDomains(sslCertInfo []types.SSLCertInfo) error {
	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	for _, cert := range sslCertInfo {
		_, err = tx.Exec(`
		DELETE from domains WHERE username = ? AND domain = ?
		`, cert.UserName, cert.Domain)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) ReadAllDomains(username string) (*[]types.SSLCertInfo, error) {
	certInfo := []types.SSLCertInfo{}

	tx, err := us.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(`
	SELECT domain, server_type, issuer, expires_in, warn_before, status, last_checked 
	FROM domains WHERE username = ?
	`, username)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		si := types.SSLCertInfo{}
		err = rows.Scan(&si.Domain, &si.ServerType, &si.Issuer, &si.ExpiresIn, &si.WarnBefore, &si.Status, &si.LastChecked)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		certInfo = append(certInfo, si)
	}

	tx.Commit()
	return &certInfo, nil
}

func (us *SqliteStore) UpdateAllDomains(username string) error {
	// TODO: update only if lastUpdatedTime is greater than certain
	// threshold

	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query(`
	SELECT domain, warn_before FROM domains WHERE username = ?
	`, username)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer rows.Close()

	var d string
	var wb int64

	for rows.Next() {
		err = rows.Scan(&d, &wb)
		if err != nil {
			tx.Rollback()
			return err
		}

		s := types.NewSSLCertInfo(d, username).WithWarnBefore(wb)
		s.Validate()

		_, err = tx.Exec(`
		UPDATE domains SET 
		server_type = ?, issuer = ?, expires_in = ?, status = ?, last_checked = ?  
		WHERE username = ? AND domain = ?
		`, s.ServerType, s.Issuer, s.ExpiresIn, s.Status, s.LastChecked, username, d)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (us *SqliteStore) CronUpdateDomains() error {
	// TODO: update only if last_checked is above
	// a certain threshold

	tx, err := us.db.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query(`
	SELECT username, domain, warn_before FROM domains
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer rows.Close()

	var u string
	var d string
	var wb int64

	for rows.Next() {
		err = rows.Scan(&u, &d, &wb)
		if err != nil {
			tx.Rollback()
			return err
		}

		s := types.NewSSLCertInfo(d, u).WithWarnBefore(wb)
		s.Validate()

		_, err = tx.Exec(`
		UPDATE domains SET 
		server_type = ?, issuer = ?, expires_in = ?, status = ?, last_checked = ?  
		WHERE username = ? AND domain = ?
		`, s.ServerType, s.Issuer, s.ExpiresIn, s.Status, s.LastChecked, u, d)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}
