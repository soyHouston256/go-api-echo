package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/soyhouston256/go-api-echo-test/authorization"

	"github.com/soyhouston256/go-api-echo-test/db"
	"github.com/soyhouston256/go-api-echo-test/handler"
	"github.com/soyhouston256/go-api-echo-test/model"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("cant load files: %v", err)
	}

	db.Connection()
	db.DB.AutoMigrate(&model.Person{}, &model.Community{})

	dbStorage := db.NewDBStorage(db.DB)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RoutePerson(e, dbStorage)
	handler.RouteLogin(e, dbStorage)

	fmt.Println("Server running on port 8081")
	err = http.ListenAndServe(":8081", e)
	if err != nil {
		fmt.Printf("cant start server: %v", err)
	}

}
