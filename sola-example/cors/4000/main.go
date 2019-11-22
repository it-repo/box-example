package main

import (
	"github.com/ddosakura/sola/v2/middleware"
	"github.com/ddosakura/sola/v2/middleware/cors"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()
	app.Use(cors.New(&cors.Option{
		Origin: func(c sola.Context) string {
			return "http://localhost:3000" // 允许 http://localhost:3000 跨域
		},
	}))

	app.Use(middleware.Static("../static", ""))
	sola.Listen("127.0.0.1:4000", app)
	sola.Keep()
}
