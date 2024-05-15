package service

import (
	"errors"
	"golang-csv/jsonFormate"
	"golang-csv/modal"
	"golang-csv/repository"
)

type UserserviceImp struct {
	User repository.UserRepo
}

func NewUserserviceImp(u repository.UserRepo) Userservice {
	return &UserserviceImp{User: u}
}

func (u *UserserviceImp) Getvalue(key string, value modal.Responder) error {
	users, err := u.User.GetAll(value)
	if err != nil {
		return err
	}
	err = jsonFormate.GetJsonFile(key, users)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserserviceImp) Setvalue(key string, value modal.Responder) error {
	fulldata, err := jsonFormate.SetJsonFile(key, value)
	if err != nil {
		return err
	}
	err = s.User.SetAll(fulldata)
	if err != nil {
		return errors.New("cannot found data in database")
	}
	return nil
}
