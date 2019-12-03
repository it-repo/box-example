package main

import (
	"fmt"
	"net/http"

	"github.com/ddosakura/sola/v2"
	router "github.com/ddosakura/sola/v2/middleware/xrouter"
)

func main() {
	app := sola.New()

	{
		r := router.New("/api/v1")
		r.Bind("/hello", hello("Hello World!"))
		{
			sub := r.Sub("/user")
			sub.Bind("/hello", hello("Hello!"))
			sub.Use(func(next sola.Handler) sola.Handler {
				return func(c sola.Context) error {
					fmt.Println("do auth")
					return next(c)
				}
			})
			sub.Bind("/info", hello("user info"))
			sub.Bind("/infox/*", hello("user infox"))
			sub.Bind("/:id", get("id"))
		}
		r.Bind("/*", error404)
		app.Use(r.Routes())
	}

	{
		r := router.New("/api/v2")
		r.Bind("/hello", hello("Hello World!"))
		{
			sub := r.Sub("/user")
			sub.Bind("/hello", hello("Hello!"))
			sub.Use(func(next sola.Handler) sola.Handler {
				return func(c sola.Context) error {
					fmt.Println("do auth")
					return next(c)
				}
			})
			sub.Bind("/info", hello("user info"))
			sub.Bind("/infox/*", hello("user infox"))
			sub.Bind("/:id", get("id"))
		}
		r.Bind("/*", error404)
		app.Use(r.Routes())
	}

	app.Use(func(sola.Handler) sola.Handler {
		return func(c sola.Context) error {
			return c.String(http.StatusOK, "no match")
		}
	})
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}

func hello(text string) sola.Handler {
	return func(c sola.Context) error {
		return c.String(http.StatusOK, text)
	}
}

func get(key string) sola.Handler {
	return func(c sola.Context) error {
		return c.String(http.StatusOK, router.Param(c, key))
	}
}

func error404(c sola.Context) error {
	return c.String(http.StatusNotFound, "404")
}
