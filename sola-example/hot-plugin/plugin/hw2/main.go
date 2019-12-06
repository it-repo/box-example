package main

import (
	"net/http"

	"github.com/ddosakura/sola/v2"
)

// export(s)
var (
	ExportHandler = map[string]sola.Handler{
		"hw2": func(c sola.Context) error {
			return c.String(http.StatusOK, "Hello World!2")
		},
	}
	ExportMiddleware = map[string]sola.Middleware{
		"hwx2": func(next sola.Handler) sola.Handler {
			return func(c sola.Context) error {
				return c.String(http.StatusOK, "Hello2-666")
			}
		},
	}
)

func main() {}
