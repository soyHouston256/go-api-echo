package handler

import (
	"encoding/json"
	"github.com/soyhouston256/go-api/authorization"
	"github.com/soyhouston256/go-api/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type login struct {
	storage Storage
}

func newLogin(storage Storage) login {
	return login{storage}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responseJSON(w, http.StatusMethodNotAllowed, messageTypeError, "method not allowed", nil)
		return
	}
	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, messageTypeError, "bad request", nil)
		return
	}
	if !isValidLogin(data, l) {
		responseJSON(w, http.StatusUnauthorized, messageTypeError, "unauthorized", nil)
		return
	}
	token, err := authorization.GenerateToken(&data)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, messageTypeError, "bad request", nil)
		return
	}

	responseJSON(w, http.StatusOK, messageTypeSuccess, "token", token)

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
