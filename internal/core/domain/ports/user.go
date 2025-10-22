package ports

import (
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/core/domain/entities"
)

type UserRepository interface {
	Save(user *entities.User) error
	FindByID(id string) (*entities.User, error)
}

type UserServices interface {
	CreateUser(cmd command.CreateUserCommand) (*entities.User, error)
	GetUser(id string) (*entities.User, error)
}
