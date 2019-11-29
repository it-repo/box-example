package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/logger"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	app := sola.New()

	app.Use(logger.New(10, func(l *logger.Log) {
		if !l.IsAction {
			log.Printf(l.Format, l.V...)
			return
		}
		fmt.Println(l.V[0], "in", l.CreateTime)
	}))
	app.Use(controller().Routes())

	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}

func controller() (r *router.Router) {
	r = router.New()

	r.BindFunc("/login", func(c sola.Context) error {
		logger.Action(c, "login")
		return c.String(http.StatusOK, "OK")
	})

	r.BindFunc("/log/:s", func(c sola.Context) error {
		s := router.Param(c, "s")
		logger.Printf(c, "%s\n", s)
		return c.String(http.StatusOK, "OK")
	})

	return
}
