package main

import (
	"fmt"
	"net/http"

	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
	_ "github.com/go-sql-driver/mysql"
	box "github.com/it-repo/box/middleware/sola-box"
	"github.com/it-repo/box/service/ac"
	"github.com/jinzhu/gorm"
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

	if db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@%s", dbUser, dbPass, dbURL)); err != nil {
		panic(err)
	} else {
		app.CacheORM("default", db)
	}

	boxRouter, requestAC := boxRoot(app)
	app.Use(boxRouter.Routes())

	r := router.New()
	acr1 := box.ACR(ac.TypeRole, ac.LogicalOR, "r2", "r3")
	r.BindFunc("/hw", requestAC(acr1, func(c sola.Context) error {
		return c.String(http.StatusOK, "Hello World! r2 | r3")
	}))
	acr2 := box.ACR(ac.TypeRole, ac.LogicalAND, "r2", "r3")
	r.BindFunc("/hw2", requestAC(acr2, func(c sola.Context) error {
		return c.String(http.StatusOK, "Hello World! r2 & r3")
	}))
	app.Use(r.Routes())

	// 监听
	sola.Listen("127.0.0.1:3000", app) // 监听 127.0.0.1:3000
	sola.Keep()                        // 固定写法，确保所有监听未结束前程序不退出
}
