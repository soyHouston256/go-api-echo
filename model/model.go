package model

import "errors"

var (
	ErrorPersonCanNotBeNil   = errors.New("person can not be nil")
	ErrPersonCanNotBeNil     = errors.New("person can not be nil")
	ErrIDPersonDoesNotExists = errors.New("id person does not exists")
)
