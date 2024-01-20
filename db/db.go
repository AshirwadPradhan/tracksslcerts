package db

import "github.com/AshirwadPradhan/tracksslcerts/types"


type UserStore interface {
	Create(types.User) error
	Update(users.User) error
	Delete(users.User) error
}