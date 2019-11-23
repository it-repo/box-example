//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js
package main

import (
	"log"
	"net/http"
	"os"

	_ "example/sola-example/native/statik"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/native"
	"github.com/ddosakura/sola/v2/middleware/router"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/afero"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	wd, _ := os.Getwd()
	bp := afero.NewBasePathFs(afero.NewOsFs(), wd)
	httpFs := afero.NewHttpFs(bp)
	fileserver := http.FileServer(httpFs.Dir("public"))

	app := sola.New()
	r := router.New()

	r.Bind("/public", native.Static("public", "/public"))
	r.BindFunc("/native", native.From(http.StripPrefix("/native", http.FileServer(http.Dir("public")))))
	r.BindFunc("/statik", native.From(http.StripPrefix("/statik", http.FileServer(statikFS))))
	r.BindFunc("/afero", native.From(http.StripPrefix("/afero", fileserver)))

	app.Use(r.Routes())
	sola.Listen("127.0.0.1:3000", app)
	sola.Keep()
}
