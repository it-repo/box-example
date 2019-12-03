package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/native"
	"github.com/ddosakura/sola/v2/middleware/proxy"
	"github.com/ddosakura/sola/v2/middleware/x/router"
)

func main() {
	app := sola.New()

	// 设置 favicon
	app.Use(proxy.Favicon("http://fanyi.bdstatic.com/static/translation/img/favicon/favicon-32x32_ca689c3.png"))

	// 和路由中间件一起使用
	r := router.New()
	r.Prefix = "/s"
	r.Bind("", native.Static("static", "/s"))
	app.Use(r.Routes())
	// 直接使用
	app.Use(native.Static(".", ""))

	sola.Listen("127.0.0.1:3000", app)

	// 自动跳转 http://127.0.0.1:3000
	bak := proxy.BackupSola("http://127.0.0.1:3000")
	sola.Listen("127.0.0.1:3001", bak)

	sola.Keep()
}
