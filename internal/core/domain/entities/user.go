package entities

import "golang-hexagonal/internal/infastructure/json/mapper"

type User struct {
	id   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		id:   "uuid",
		Name: name,
	}
}

func NewUserFromDAO(user *mapper.UserReconsitutionData) *User {
	return &User{
		id:   user.ID,
		Name: user.FullName,
	}
}

func (u *User) IsNameValid() bool {
	return u.Name != ""
}
