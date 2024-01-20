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
		_, err = tx.Exec(`INSERT INTO domains 
		(username, domain, server_type, issuer, expires_in, warn_before, status, last_checked)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			cert.UserName, cert.Domain, cert.ServerType, cert.Issuer, cert.ExpiresIn, cert.WarnBefore, cert.Status, cert.LastChecked)
		
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

	rows, err := tx.Query("select username, email, account_type from users where username = ?", username)
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
