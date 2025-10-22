package utils

import "github.com/google/uuid"

func UserID() (id string) {

	uuidRaw := uuid.New()
	uuidStr := uuidRaw.String()

	return uuidStr
}
