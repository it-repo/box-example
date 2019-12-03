package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/auth"
	"github.com/ddosakura/sola/v2/middleware/x/router"
	box "github.com/it-repo/box/middleware/sola-box"
	"github.com/it-repo/box/service/ac"
	"github.com/spf13/viper"
)

func boxRoot(app *sola.Sola) (*router.Router, box.RequestAC) {
	solaAuthKey := viper.GetString("sola.auth.key")
	logger, logList := box.Logger(10, app.DefaultORM())
	app.Use(logger)

	acInit, acRouter, requestAC := box.AC(app.DefaultORM(), solaAuthKey)

	r := router.New()
	r.Prefix = "/api/box"
	{
		acRouter.Prefix = r.Prefix + "/user"
		r.Bind("/user", sola.Merge(acInit, acRouter.Routes()))

		routeRouter := box.Route(app.DefaultORM())
		routeRouter.Prefix = r.Prefix
		_auth := auth.Auth(auth.AuthJWT, []byte(solaAuthKey))
		r.Bind("/routes", auth.New(_auth, nil, routeRouter.Routes()))

		acrLogList := box.ACR(ac.TypeRole, ac.LogicalOR, "admin")
		r.BindFunc("/sys/log", requestAC(acrLogList, logList))
	}

	return r, requestAC
}
