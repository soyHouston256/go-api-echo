package handler

import "github.com/soyhouston256/go-api-echo-test/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetAll() (model.Persons, error)
	GetByID(ID int) (*model.Person, error)
	FindUserByEmail(email string) (*model.Person, error)
}
