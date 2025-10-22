package mapper

import "golang-hexagonal/internal/infastructure/json/models"

type UserReconsitutionData struct {
	ID       string
	FullName string
}

func DaoToDomain(user *models.User) *UserReconsitutionData {
	return &UserReconsitutionData{
		ID:       user.Id,
		FullName: user.FullName,
	}
}
