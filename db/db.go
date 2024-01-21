package db

import "github.com/AshirwadPradhan/tracksslcerts/types"

// TODO: check if pointers can be used to pass all around

type UserStorer interface {
	CreateUser(*types.User) error
	ReadUser(string) (*types.User, error)
	DeleteUser(*types.User) error
	UserUpdater
}

type UserUpdater interface {
	UpdateUserUsername(string, *types.User) error
	UpdateUserPassword(string, *types.User) error
	UpdateUserAccountType(string, *types.User) error
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
