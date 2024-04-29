package db

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/soyhouston256/go-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	dsn = "host=localhost user=user password=admin dbname=pichangas port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

type Storage struct {
	DB *gorm.DB
}

func NewDBStorage(db *gorm.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func Connection() {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("cant open db connection: %v", err)
		panic(err)
	}
	fmt.Println("connected to db")
}

func (s *Storage) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	hashedPassword, err := HashPassword(person.Password)
	if err != nil {
		return err
	}
	person.Password = string(hashedPassword)

	result := s.DB.Create(person)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Storage) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	result := s.DB.Model(&model.Person{}).Where("id = ?", ID).Updates(person)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return nil
}

func (s *Storage) Delete(ID int) error {
	result := s.DB.Delete(&model.Person{}, ID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return nil
}

func (s *Storage) GetAll() (model.Persons, error) {
	var persons model.Persons
	result := s.DB.Preload("Communities").Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}
	return persons, nil
}

func (s *Storage) GetByID(ID int) (*model.Person, error) {
	var person model.Person
	result := s.DB.Preload("Communities").First(&person, ID)
	if result.Error != nil {
		return &person, result.Error
	}
	if result.RowsAffected == 0 {
		return &person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return &person, nil
}

func (s *Storage) FindUserByEmail(email string) (*model.Person, error) {
	var person model.Person
	result := s.DB.Where("email = ?", email).First(&person)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &person, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
