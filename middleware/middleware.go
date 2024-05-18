package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/soyhouston256/go-api-echo-test/authorization"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		log.Printf("Logging petition %q, metohd: %q", r.URL.Path, r.Method)
		f(w, r)
		log.Println(time.Since(t0))
	}
}

func Authenticated(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			return ctx.JSON(http.StatusForbidden, map[string]string{"message": "forbidden"})
		}
		return f(ctx)
	}
}
