package main

import (
	"fmt"
	"github.com/soyhouston256/go-api/authorization"
	"log"
	"net/http"

	"github.com/soyhouston256/go-api/db"
	"github.com/soyhouston256/go-api/handler"
	"github.com/soyhouston256/go-api/model"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("cant load files: %v", err)
	}

	db.Connection()
	db.DB.AutoMigrate(&model.Person{}, &model.Community{})

	dbStorage := db.NewDBStorage(db.DB)

	mux := http.NewServeMux()
	handler.RoutePerson(mux, dbStorage)
	handler.RouteLogin(mux, dbStorage)

	fmt.Println("Server running on port 8081")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Printf("cant start server: %v", err)
	}

}
