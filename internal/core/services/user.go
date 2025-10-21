package services

import (
	"errors"
	"golang-hexagonal/internal/core/command"
	"golang-hexagonal/internal/core/domain/entities"
	"golang-hexagonal/internal/core/domain/ports"
	cmdMapper "golang-hexagonal/internal/core/mapper"
)

type UserServices struct {
	repo ports.UserRepository
}

/*
เข้าใจว่าเรายัด CreateUser(name string) (*entities.User, error)
เข้าไปใน UserServices แล้วพอขา return ไปที่ ports.UserServices
มันตรวจสอบว่ามีสัญญาที่ตรงกันไหม

(u *UserServices) คือ method หรือ การยัดเข้าไป
*/
func NewUserServices(
	repo ports.UserRepository, /*repo ports.UserRepository parameter นี้ไว้รับการเชื่อมต่อจาก DB *รับมาจาก main.go*/
) ports.UserServices {
	return &UserServices{
		repo: repo, /*เอา param ที่รับเข้ามากำหนดให้ที่ reference ไว้ใน UserServices ทำให้เข้าถึง repository ได้เลย*/
	}
}

func (u *UserServices) CreateUser(cmd command.CreateUserCommand) (*entities.User, error) {

	newUser := cmdMapper.ToDomain(&cmd)

	if !newUser.IsNameValid() {
		return nil, errors.New("user name validation failed")
	}

	if existUser, _ := u.repo.FindByName(newUser.Name); existUser != nil {
		return nil, errors.New("user already exists")
	}

	if err := u.repo.Save(newUser); err != nil {
		return nil, errors.New("user already exists")
	}

	return nil, nil
}

/*

! Tips !

การเขียน return error

แบบที่ 1 (Limited Scope)

	if v , err := .... ; err != nil {
		return err
	}

แบบที่ 2 (Wider Scope)

	err := ....
	if err != nil {
		return err
	}

	err = .... <- err ตรงนี้ยังใช้ได้

	แบบที่ 1 จะเหมาะกว่าในเนื่องของการกำหนด scope เพราะมันช่วยจำกัด Scope ของตัวแปร err
	ซึ่งเป็นสิ่งที่ดีที่สุดในทางปฏิบัติเพื่อหลีกเลี่ยงข้อผิดพลาดจากตัวแปรซ้ำซ้อน

	แบบที่ 2 จะสามารถใช้ err ได้ต่อจากการประกาศครั้งแรก หมายความมันยังใช้ได้นอก scope if
*/
