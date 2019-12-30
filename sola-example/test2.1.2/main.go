package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ddosakura/sola/v2/middleware/router"

	"github.com/ddosakura/sola/v2"
)

func hw(c sola.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

var m = sola.M(func(c sola.C, next sola.H) error {
	fmt.Println("h2m")
	// return c.String(http.StatusOK, "h2m")
	return next(c)
}).M()

func main() {
	sola.DefaultApp.Dev()
	hLog := sola.Handler(func(c sola.Context) error {
		log.Println("Hello World!")
		return nil
	})
	hError := sola.Handler(func(c sola.Context) error {
		return sola.Error("sakura", errors.New("test"))
	})
	r := router.New(nil)
	r.Bind("/hell", hw)
	r.Bind("/hello", hw, hLog.M())
	r.Bind("/helloo", hw)
	r.Bind("/error", hError)
	r.Bind("/h2m", m.H())
	sola.Use(r.Routes())
	sola.ListenKeep("127.0.0.1:3000")
}
