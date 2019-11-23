package main

import (
	"github.com/ddosakura/sola/v2/middleware/cors"
	"github.com/ddosakura/sola/v2/middleware/native"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()
	app.Use(cors.New(nil)) // 允许跨域

	app.Use(native.Static("../static", ""))
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
