package repository

import "golang-csv/modal"

type UserRepo interface {
	GetAll(value modal.Responder) (interface{}, error)
	SetAll(user modal.Responder) error
}
