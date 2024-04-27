package main

import (
	"fmt"
	"net/http"

	"github.com/soyhouston256/go-api/db"
	"github.com/soyhouston256/go-api/handler"
	"github.com/soyhouston256/go-api/model"
)

func main() {
	db.Connection()
	db.DB.AutoMigrate(&model.Person{}, &model.Community{})

	dbStorage := db.NewDBStorage(db.DB)

	mux := http.NewServeMux()
	handler.RoutePerson(mux, dbStorage)
	fmt.Println("Server running on port 8081")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Printf("cant start server: %v", err)
	}

}
