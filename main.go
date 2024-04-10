package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", greating)
	http.ListenAndServe(":8080", nil)
}

func greating(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
