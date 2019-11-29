package main

import (
	"example/sola-example/restful-api/controller"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()
	r := controller.UserC()
	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
