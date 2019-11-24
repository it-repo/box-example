package main

import (
	"os"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/proxy"
	"github.com/spf13/afero"
)

func main() {
	app := sola.New()
	app.Dev()

	// 设置 favicon
	app.Use(proxy.Favicon("http://fanyi.bdstatic.com/static/translation/img/favicon/favicon-32x32_ca689c3.png"))

	app.Use(proxy.New(`function handle()
	if (URL == "/hw")
	then
		return 200, "Hello World!"
	elseif (URL == "/hw2")
	then
		return 301, "/hw"
	end
end`))

	wd, _ := os.Getwd()
	bp := afero.NewBasePathFs(afero.NewOsFs(), wd)
	script, _ := afero.ReadFile(bp, "main.lua")
	app.Use(proxy.New(string(script)))

	sola.Listen("127.0.0.1:3000", app)

	// 自动跳转 http://127.0.0.1:3000
	bak := proxy.BackupSola("http://127.0.0.1:3000")
	sola.Listen("127.0.0.1:3001", bak)

	sola.Keep()
}
