package middleware

import (
	"log"
	"net/http"
	"time"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		log.Printf("Logging petition %q, metohd: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Println(time.Since(t0))
	}
}

func Authenticated(f func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret-token-here" {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"message": "forbidden"}`))
}
