package mapper

import (
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/presentations/dto"
)

func ToCommand(req *dto.CraeteUserRequest) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		FullName: req.FirstName + " " + req.LastName,
	}
}
