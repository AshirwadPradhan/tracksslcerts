package types

type User struct {
	UserName       string
	Email          string
	HashedPassword string
	AccountType    string
}

func NewUser(userName string, email string, password string) *User {
	return &User{
		UserName:       userName,
		Email:          email,
		HashedPassword: password,
	}
}

// func (u *User) AddDomainToTrack(domain string) {
// 	sci := NewSSLCertInfo(domain)
// 	sci.Validate()
// 	u.TrackedDomains = append(u.TrackedDomains, *sci)
// }

// func (u *User) RemoveTrackedDomain(domain string) {
// 	idx := 0
// 	for i, d := range u.TrackedDomains {
// 		if d.Domain == domain {
// 			idx = i
// 			break
// 		}
// 	}
// 	before := u.TrackedDomains[:idx]
// 	u.TrackedDomains = append(before, u.TrackedDomains[idx+1:]...)
// }
