package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/soyhouston256/go-api-echo-test/middleware"
)

func RoutePerson(e *echo.Echo, storage Storage) {
	p := newPerson(storage)
	persons := e.Group("/v1/persons")
	persons.Use(middleware.Authenticated)
	persons.POST("", p.create)
	persons.PUT("/:id", p.update)
	persons.DELETE("/:id", p.delete)
	persons.GET("/:id", p.getById)
	persons.GET("", p.getAll)
}

func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
