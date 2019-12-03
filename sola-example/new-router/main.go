package main

import (
	"fmt"
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func lookMeta(next sola.Handler) sola.Handler {
	return func(c sola.Context) error {
		fmt.Println(c.Get(router.CtxMeta))
		return next(c)
	}
}

func main() {
	app := sola.New()

	{
		r := router.New(&router.Option{
			Pattern: "/api",
		})
		app.Use(r.Routes())
	}

	app.Use(lookMeta) // test shadow context & back to origin

	{
		r := router.New(&router.Option{
			Pattern: "/api/v1",
		})
		r.Bind("/hello", hello("Hello World!"))
		{
			sub := r.Sub(&router.Option{
				Pattern: "/user",
			})
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
		r := router.New(&router.Option{
			Pattern: "/api/v2",
		})
		r.Bind("/hello", hello("Hello World!"))
		{
			sub := r.Sub(&router.Option{
				Pattern:     "/user",
				UseNotFound: true,
			})
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
