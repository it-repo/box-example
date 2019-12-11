package main

import (
	"github.com/ddosakura/sola/v2/middleware/proxy"

	"github.com/ddosakura/sola/v2"
)

func main() {
	app := sola.New()
	app.Dev()

	app.Use(proxy.New(`function handle()
	proxy("https://blog.moyinzi.top")
	-- proxy("https://github.com/ddosakura")
	return 0
end`))

	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
