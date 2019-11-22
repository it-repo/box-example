package main

import (
	"github.com/ddosakura/sola/v2/middleware"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()

	app.Use(middleware.Static("../static", ""))
	sola.Listen("127.0.0.1:2000", app)
	sola.Keep()
}
