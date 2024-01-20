package db

import "github.com/AshirwadPradhan/tracksslcerts/types"


type UserStore interface {
	Create(types.User) error
	Update(types.User) error
	Delete(types.User) error
}