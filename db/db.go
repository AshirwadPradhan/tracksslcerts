package db

import "github.com/AshirwadPradhan/tracksslcerts/types"

type UserStorer interface {
	Create(types.User) error
	Read(string) (types.User, error)
	Delete(types.User) error
	UserUpdater
}

type UserUpdater interface {
	UpdateUsername(string) error
	UpdatePassword(string) error
	UpdateAccountType(string) error
}

type DomainStorer interface {
	Add(string) error
	Read(string) (types.SSLCertInfo, error)
	Delete(string) error
}
