package models

import (
	"golang-hexagonal/utils"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
}

func NewUser(user User) *User {
	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	return &User{
		Id:       utils.UserID(),
		FullName: firstName + " " + lastName,
	}
}
