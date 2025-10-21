package ports

import (
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/core/domain/entities"
)

type UserRepository interface {
	Save(user *entities.User) error
	FindByName(name string) (*entities.User, error)
}

type UserServices interface {
	CreateUser(cmd command.CreateUserCommand) (*entities.User, error)
}
