// db_test.go
package db

import (
	"fmt"
	"testing"

	"github.com/soyhouston256/go-api-echo-test/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.Person{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestStorage_Create(t *testing.T) {

	table := []struct {
		name      string
		person    *model.Person
		wantError error
	}{
		{
			name:      "Test Table",
			person:    &model.Person{Name: "Test User", Age: 20, Email: "test@example.com", Password: "password"},
			wantError: nil,
		},
		{
			name:      "Test error",
			person:    nil,
			wantError: model.ErrPersonCanNotBeNil,
		},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			db := setupTestDB(t)
			storage := NewDBStorage(db)
			gotError := storage.Create(v.person)

			fmt.Printf("gotError: %v\n", gotError)

			if gotError != v.wantError {
				t.Errorf("expected error %v, got %v", v.wantError, gotError)
			}

		})
	}
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person := &model.Person{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}

	err := storage.Create(person)
	if err != nil {
		t.Errorf("failed to create person: %v", err)
	}

	var result model.Person
	err = db.First(&result, "email = ?", person.Email).Error
	if err != nil {
		t.Errorf("failed to find created person: %v", err)
	}
	if result.Name != person.Name || result.Email != person.Email {
		t.Errorf("created person does not match: got %v, want %v", result, person)
	}
}

func TestStorage_Update(t *testing.T) {
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person := &model.Person{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(person)

	updatedPerson := &model.Person{
		Name:     "Updated User",
		Email:    "updated@example.com",
		Password: "newpassword",
	}

	err := storage.Update(int(person.ID), updatedPerson)
	if err != nil {
		t.Errorf("failed to update person: %v", err)
	}

	var result model.Person
	err = db.First(&result, person.ID).Error
	if err != nil {
		t.Errorf("failed to find updated person: %v", err)
	}
	if result.Name != updatedPerson.Name || result.Email != updatedPerson.Email {
		t.Errorf("updated person does not match: got %v, want %v", result, updatedPerson)
	}
}

func TestStorage_Delete(t *testing.T) {
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person := &model.Person{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(person)

	err := storage.Delete(int(person.ID))
	if err != nil {
		t.Errorf("failed to delete person: %v", err)
	}

	var result model.Person
	err = db.First(&result, person.ID).Error
	if err == nil {
		t.Errorf("person should have been deleted, found: %v", result)
	}
}

func TestStorage_GetAll(t *testing.T) {
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person1 := &model.Person{
		Name:     "User One",
		Email:    "one@example.com",
		Password: "password",
	}
	person2 := &model.Person{
		Name:     "User Two",
		Email:    "two@example.com",
		Password: "password",
	}
	db.Create(person1)
	db.Create(person2)

	persons, err := storage.GetAll()
	if err != nil {
		t.Errorf("failed to get all persons: %v", err)
	}
	if len(persons) != 2 {
		t.Errorf("unexpected number of persons: got %v, want 2", len(persons))
	}
}

func TestStorage_GetByID(t *testing.T) {
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person := &model.Person{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(person)

	result, err := storage.GetByID(int(person.ID))
	if err != nil {
		t.Errorf("failed to get person by ID: %v", err)
	}
	if result.Name != person.Name || result.Email != person.Email {
		t.Errorf("person does not match: got %v, want %v", result, person)
	}
}

func TestStorage_FindUserByEmail(t *testing.T) {
	db := setupTestDB(t)
	storage := NewDBStorage(db)

	person := &model.Person{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	db.Create(person)

	result, err := storage.FindUserByEmail(person.Email)
	if err != nil {
		t.Errorf("failed to find user by email: %v", err)
	}
	if result.Name != person.Name || result.Email != person.Email {
		t.Errorf("person does not match: got %v, want %v", result, person)
	}
}
