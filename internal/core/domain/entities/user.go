package entities

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

func (u *User) IsNameValid() bool {
	return u.Name != ""
}
