package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/soyhouston256/go-api-echo-test/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p person) create(ctx echo.Context) error {

	data := model.Person{}
	err := ctx.Bind(&data)
	if err != nil {
		resp := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	err = p.storage.Create(&data)
	if err != nil {
		resp := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusInternalServerError, resp)
	}
	resp := newResponse(messageTypeSuccess, "created", nil)
	return ctx.JSON(http.StatusCreated, resp)
}

func (p person) update(ctx echo.Context) error {

	ID, error := strconv.Atoi(ctx.Param("id"))
	if error != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err := ctx.Bind(&data)
	if err != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(messageTypeSuccess, "updated", nil)
	return ctx.JSON(http.StatusOK, response)

}

func (p person) delete(ctx echo.Context) error {

	ID, error := strconv.Atoi(ctx.Param("id"))
	if error != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	err := p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(messageTypeError, "person not found", nil)
		return ctx.JSON(http.StatusNotFound, response)
	}
	if err != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(messageTypeSuccess, "deleted", nil)
	return ctx.JSON(http.StatusOK, response)
}

func (p person) getById(ctx echo.Context) error {

	ID, error := strconv.Atoi(ctx.Param("id"))
	if error != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(messageTypeError, "person not found", nil)
		return ctx.JSON(http.StatusNotFound, response)
	}
	response := newResponse(messageTypeSuccess, "success", data)
	return ctx.JSON(http.StatusOK, response)

}

func (p person) getAll(ctx echo.Context) error {

	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(messageTypeSuccess, "success", data)
	return ctx.JSON(http.StatusOK, response)
}
