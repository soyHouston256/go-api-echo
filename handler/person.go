package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/soyhouston256/go-api/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responseJSON(w, http.StatusMethodNotAllowed, messageTypeError, "method not allowed", nil)
		return
	}
	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, messageTypeError, "bad request", nil)
		return
	}
	err = p.storage.Create(&data)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, messageTypeError, "bad request", nil)
		return
	}
	responseJSON(w, http.StatusCreated, messageTypeSuccess, "created", nil)
}

func (p person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		responseJSON(w, http.StatusMethodNotAllowed, messageTypeError, "method not allowed", nil)
		return
	}

	ID, error := strconv.Atoi(r.URL.Query().Get("id"))
	if error != nil {
		responseJSON(w, http.StatusBadRequest, messageTypeError, "bad request", nil)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, messageTypeError, "bad request", nil)
		return
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, messageTypeError, "bad request", nil)
		return
	}
	responseJSON(w, http.StatusOK, messageTypeSuccess, "updated", nil)

}

func (p person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		responseJSON(w, http.StatusMethodNotAllowed, messageTypeError, "method not allowed", nil)
		return
	}

	ID, error := strconv.Atoi(r.URL.Query().Get("id"))
	if error != nil {
		responseJSON(w, http.StatusBadRequest, messageTypeError, "bad request", nil)
		return
	}

	err := p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		responseJSON(w, http.StatusNotFound, messageTypeError, "person not found", nil)
		return
	}
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, messageTypeError, "bad request", nil)
		return
	}
	responseJSON(w, http.StatusOK, messageTypeSuccess, "deleted", nil)

}

func (p person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responseJSON(w, http.StatusMethodNotAllowed, messageTypeError, "method not allowed", nil)
		return
	}
	data, err := p.storage.GetAll()
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, messageTypeError, "bad request", nil)
		return
	}
	responseJSON(w, http.StatusOK, messageTypeSuccess, "success", &data)
}
