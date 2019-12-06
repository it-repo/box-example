package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ddosakura/sola/v2"
	linux "github.com/ddosakura/sola/v2/extension/sola-linux"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	app := sola.New()
	r := router.New(&router.Option{
		UseNotFound: true,
	})

	r.Bind("/:t", func(c sola.Context) error {
		t := router.Param(c, "t")
		n, _ := strconv.Atoi(t)
		time.Sleep(time.Duration(n * int(time.Second)))
		return c.String(http.StatusOK, "sleep "+t)
	})

	app.Use(r.Routes())
	linux.Listen("127.0.0.1:3000", app)
	linux.Keep()
}
