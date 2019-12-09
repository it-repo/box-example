package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/extension/hot"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	h, err := hot.New(&hot.Option{
		Init: []string{
			"./plugin/hw.so",
			"./plugin/hw2",
		},
		Watch: []string{"./plugin/hw2"},
	})
	if err != nil {
		panic(err)
	}

	app := sola.New()
	app.Dev()
	h.Used(app)
	r := router.New(nil)
	r.Bind("/info", func(c sola.Context) error {
		h := h.Handler("hw")
		return h(c)
	})
	r.Bind("/hello", h.Handler("hw"))
	app.Use(r.Routes())
	app.Use(h.Middleware("hwx2"))
	sola.Listen("127.0.0.1:3000", app)
	h.Scan()
}
