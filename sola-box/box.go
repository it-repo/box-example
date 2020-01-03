package main

import (
	"github.com/ddosakura/sola/v2"
	"github.com/ddosakura/sola/v2/middleware/router"
	box "github.com/it-repo/box/middleware/sola-box"
	"github.com/it-repo/box/service/ac"
	"github.com/spf13/viper"
)

func boxRoot(app *sola.Sola, r *router.Router) (sola.Middleware, box.ACRequest) {
	key := viper.GetString("sola.auth.key")
	logger, logList := box.Logger(10, app.DefaultORM())
	app.Use(logger)

	jwtAuth, acRequest := box.AC(app.DefaultORM(), key, r.Sub(&router.Option{
		Pattern: "/user",
	}))
	r.Use(jwtAuth)
	box.Route(app.DefaultORM(), r, acRequest)

	acrLogList := box.ACR(ac.TypeRole, ac.LogicalOR, "admin")
	r.Bind("/sys/log", acRequest(acrLogList, logList))

	return jwtAuth, acRequest
}
