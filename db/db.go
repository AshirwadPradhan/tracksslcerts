package db

import "github.com/AshirwadPradhan/tracksslcerts/types"

type UserStorer interface {
	Create(types.User) error
	Read(string) (types.User, error)
	Update(types.User) error
	Delete(types.User) error
}
