package handler

import "net/http"

func RoutePerson(mux *http.ServeMux, storage Storage) {
	p := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", p.create)

	mux.HandleFunc("/v1/persons/get-all", p.getAll)

	mux.HandleFunc("/v1/persons/update", p.update)

	mux.HandleFunc("/v1/persons/delete", p.delete)
}
