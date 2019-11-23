package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/native"
)

func main() {
	app := sola.New()

	app.Use(native.Static("../static", ""))
	sola.Listen("127.0.0.1:2000", app)
	sola.Keep()
}
