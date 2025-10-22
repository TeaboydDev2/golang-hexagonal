package repository

import (
	"encoding/json"
	"fmt"
	"golang-hexagonal/internal/core/domain/entities"
	"golang-hexagonal/internal/core/domain/ports"
	"golang-hexagonal/internal/infastructure/json/mapper"
	"golang-hexagonal/internal/infastructure/json/models"
	"os"
)

const fileName = "./../infastructure/json/collection/user.json"

type JsonUserRepository struct{}

func NewJsonUserRepository() ports.UserRepository {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if err := os.MkdirAll("collections", 0755); err != nil {
			fmt.Errorf("cannot file json file")
		}
		if err := os.WriteFile(fileName, []byte("[]"), 0644); err != nil {
			fmt.Errorf("cannot file json file")
		}
	}

	return &JsonUserRepository{}
}

func (u *JsonUserRepository) loadAllDAO() ([]*models.User, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var daos []*models.User
	if err := json.Unmarshal(data, &daos); err != nil {
		return nil, err
	}
	return daos, nil
}

func (u *JsonUserRepository) Save(user *entities.User) error {
	return nil
}

func (u *JsonUserRepository) FindByID(id string) (*entities.User, error) {

	daos, err := u.loadAllDAO()
	if err != nil {
		return nil, err
	}

	for _, dao := range daos {
		if dao.Id == id {

			entity := mapper.DaoToDomain(dao)

			return entities.NewUserFromDAO(entity), nil
		}
	}

	return nil, nil // ไม่พบผู้ใช้
}

// func JsonUser() error {

// 	const count = 100
// 	const fileName = "./../infastructure/json/collection/user.json"

// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	if _, err := file.WriteString("[\n"); err != nil {
// 		return err
// 	}

// 	for i := 0; i <= count; i++ {

// 		user := models.NewUser(models.User{})

// 		userJson, err := json.MarshalIndent(user, " ", " ")
// 		if err != nil {
// 			return err
// 		}

// 		if _, err := file.Write(userJson); err != nil {
// 			return err
// 		}

// 		if i < count {
// 			if _, err := file.WriteString(",\n"); err != nil {
// 				return err
// 			} else {
// 				if _, err := file.WriteString("\n"); err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}

// 	if _, err := file.WriteString("]"); err != nil {
// 		return err
// 	}

// 	return nil
// }
