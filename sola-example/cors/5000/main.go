package main

import (
	"net/http"

	"github.com/ddosakura/sola/v2/middleware/cors"
	"github.com/ddosakura/sola/v2/middleware/native"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()
	app.Use(cors.New(&cors.Option{
		AllowMethods: []string{http.MethodOptions, http.MethodPost}, // 允许 OPTIONS,POST 跨域
	}))

	app.Use(native.Static("../static", ""))
	sola.Listen("127.0.0.1:5000", app)
	sola.Keep()
}
