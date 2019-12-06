package main

import (
	"net/http"

	"github.com/ddosakura/sola/v2"
)

// export(s)
var (
	ExportHandler = map[string]sola.Handler{
		"hw": func(c sola.Context) error {
			return c.String(http.StatusOK, "Hello World!")
		},
	}
	ExportMiddleware = map[string]sola.Middleware{
		"hwx": func(next sola.Handler) sola.Handler {
			return func(c sola.Context) error {
				return c.String(http.StatusOK, "Hello")
			}
		},
	}
)

func main() {}
