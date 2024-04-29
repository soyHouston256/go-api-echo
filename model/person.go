package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name        string
	Age         int
	Email       string
	PhoneNumber string
	Password    string
	Communities []Community
}

type Community struct {
	gorm.Model
	Name     string
	PersonID uint
}

type Persons []Person
