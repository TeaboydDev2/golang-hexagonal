package mapper

import (
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/core/domain/entities"
)

func ToDomain(user *command.CreateUserCommand) *entities.User {
	return entities.NewUser(user.FullName)
}
