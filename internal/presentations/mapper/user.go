package mapper

import (
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/core/domain/entities"
	"golang-hexagonal/internal/presentations/dto"
)

func ToCommand(req *dto.CraeteUserRequest) *command.CreateUserCommand {
	return &command.CreateUserCommand{
		FullName: req.FirstName + " " + req.LastName,
	}
}

func ToHandler(entities *entities.User) *dto.UserResponse {
	return &dto.UserResponse{
		FullName: entities.Name,
	}
}
