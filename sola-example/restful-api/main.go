package main

import (
	"example/sola-example/restful-api/controller"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
)

func main() {
	app := sola.New()
	r := controller.UserC(nil)

	root := router.New(&router.Option{
		UseNotFound: true,
	})
	root.Sub(&router.Option{
		Pattern: "/sub",
	}).Use(r.Routes())

	app.Use(r.Routes(), root.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
