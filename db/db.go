package db

import "github.com/AshirwadPradhan/tracksslcerts/types"

type UserStorer interface {
	CreateUser(types.User) error
	ReadUser(string) (types.User, error)
	DeleteUser(types.User) error
	UserUpdater
}

type UserUpdater interface {
	UpdateUserUsername(string) error
	UpdateUserPassword(string) error
	UpdateUserAccountType(string) error
}

type DomainStorer interface {
	AddDomainsToTrack([]types.SSLCertInfo) error
	ReadAllDomains(string) ([]types.SSLCertInfo, error)
	DeleteDomains([]types.SSLCertInfo) error
	DomainUpdater
}

type DomainUpdater interface {
	UpdateAllDomains(username string) error
	CronUpdateDomains() error
}
