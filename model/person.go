package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name        string
	Age         int
	Communities []Community
}

type Community struct {
	gorm.Model
	Name     string
	PersonID uint
}

type Persons []Person
