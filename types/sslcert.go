package types

import (
	"crypto/tls"
	"time"

	"github.com/AshirwadPradhan/tracksslcerts/helpers"
)

type CertStatus int

const (
	CertHealthy CertStatus = iota
	CertAboutTime
	CertUnhealthy
	CertInvalid
)

type SSLCertInfo struct {
	Domain     string
	Conn       *tls.Conn
	ServerType string
	Issuer     string

	// calculates how much time after the certificate will expire
	// stores in days
	ExpiresIn int64

	// No of days before the ceritificate expiry when emails reminders
	// is sent
	WarnBefore int64

	// Stores the status of SSL certificate
	// Healthy, if ExpiresIn is more that no. of warning days
	// AboutTime, if ExpiresIn is between 0 and warning days
	// Invalid, if SSL certificate is invalid
	Status CertStatus
	
	// Latest timestamp when Validate was called
	LastChecked time.Time
}

func NewSSLCertInfo(domain string) *SSLCertInfo {
	return &SSLCertInfo{
		Domain:     domain,
		WarnBefore: 15,
	}
}

func (s *SSLCertInfo) WithWarnBefore(warnDaysBefore int64) *SSLCertInfo {
	return &SSLCertInfo{
		WarnBefore: warnDaysBefore,
	}
}

func (s *SSLCertInfo) Validate() {
	s.LastChecked = time.Now()
	if err := s.checkValidSSL(); err != nil {
		s.Status = CertInvalid
		return
	}

	s.populateCertInfo()

	if s.ExpiresIn > s.WarnBefore {
		s.Status = CertHealthy
	} else if s.ExpiresIn >= 0 && s.ExpiresIn <= s.WarnBefore {
		s.Status = CertAboutTime
	} else if s.ExpiresIn < 0 {
		s.Status = CertUnhealthy
	}

	s.Conn.Close()
}

func (s *SSLCertInfo) populateCertInfo() {
	cs := s.Conn.ConnectionState()
	s.Issuer = cs.PeerCertificates[0].Issuer.String()
	s.ServerType = cs.ServerName
	s.ExpiresIn = helpers.CalcDaysDiff(cs.PeerCertificates[0].NotAfter, time.Now())
}

func (s *SSLCertInfo) checkValidSSL() error {
	var err error
	s.Conn, err = tls.Dial("tcp", s.Domain+":443", nil)
	if err != nil {
		return err
	}

	if err = s.Conn.VerifyHostname(s.Domain); err != nil {
		return err
	}
	return nil
}