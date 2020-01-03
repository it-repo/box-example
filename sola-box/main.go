//go:generate statik -f -src=./public -include=*.jpg,*.png,*.gif,*.ico,*.txt,*.html,*.css,*.js,*.woff,*.ttf
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "example/sola-box/statik"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/native"
	"github.com/ddosakura/sola/v2/middleware/router"
	_ "github.com/go-sql-driver/mysql"
	box "github.com/it-repo/box/middleware/sola-box"
	"github.com/it-repo/box/service/ac"
	"github.com/jinzhu/gorm"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/viper"
)

func main() {
	app := sola.New()
	app.LoadConfig()
	dbDriver := viper.GetString("db.driver")
	dbURL := viper.GetString("db.url")
	dbUser := viper.GetString("db.user")
	dbPass := viper.GetString("db.pass")
	// TODO: boxSalt := viper.GetString("box.ac.salt")

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	if db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@%s", dbUser, dbPass, dbURL)); err != nil {
		panic(err)
	} else {
		app.CacheORM("default", db)
	}

	r := router.New(nil)

	{
		base := viper.GetString("box.base")
		jwtAuth, acRequest := boxRoot(app, r.Sub(&router.Option{
			Pattern:     base + "/api/box",
			UseNotFound: true,
		}))

		r.Use(jwtAuth) // 加载用户登录信息

		acr1 := box.ACR(ac.TypeRole, ac.LogicalOR, "r2", "r3")
		r.Bind("/hw", acRequest(acr1, func(c sola.Context) error {
			return c.String(http.StatusOK, "Hello World! r2 | r3")
		}))
		acr2 := box.ACR(ac.TypeRole, ac.LogicalAND, "r2", "r3")
		r.Bind("/hw2", acRequest(acr2, func(c sola.Context) error {
			return c.String(http.StatusOK, "Hello World! r2 & r3")
		}))

		r.Bind("GET /*", native.From(http.StripPrefix("", http.FileServer(statikFS))))
	}

	app.Use(r.Routes())

	host := viper.GetString("box.host")
	port := viper.GetString("box.port")
	host += ":" + port
	sola.Listen(host, app)
	log.Println("listen", host)
	sola.Keep()
}
