package types

type User struct {
	UserName string
	Email    string
	// TODO: Add salt to the password
	HashedPassword string
	AccountType    string
}

func NewUser(userName string, email string, hpass string) *User {
	return &User{
		UserName:       userName,
		Email:          email,
		HashedPassword: hpass,
		AccountType:    "Free",
	}
}
