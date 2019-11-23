package main

import (
	"errors"
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	app := sola.New()
	app.SetHandler(http.StatusNotFound, func(c sola.Context) error {
		if c.Request().URL.String() == "/err" {
			return errors.New("/err")
		}
		return c.String(http.StatusNotFound, "404")
	})
	app.ErrorHandler = func(e error, c sola.Context) {
		c.String(http.StatusInternalServerError, e.Error())
	}

	r := router.New()
	r.BindFunc("/hw", func(c sola.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
