package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/soyhouston256/go-api-echo-test/authorization"
	"github.com/soyhouston256/go-api-echo-test/model"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	storage Storage
}

func newLogin(storage Storage) login {
	return login{storage}
}

func (l *login) login(ctx echo.Context) error {

	data := model.Login{}
	err := ctx.Bind(&data)
	if err != nil {
		resp := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	if !isValidLogin(data, l) {
		resp := newResponse(messageTypeError, "unauthorized", nil)
		return ctx.JSON(http.StatusUnauthorized, resp)
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(messageTypeError, "bad request", nil)
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(messageTypeSuccess, "token", token)
	return ctx.JSON(http.StatusOK, resp)

}

func isValidLogin(data model.Login, l *login) bool {

	personInfo, err := l.storage.FindUserByEmail(data.Email)
	if err != nil {
		return false
	}

	if CheckPasswordHash(data.Password, personInfo.Password) {
		return true
	}

	return false
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
